terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.do_token
}

# Use a data source to fetch existing VPC if it exists
data "digitalocean_vpc" "playground_vpc" {
  name = "playground-network-${var.environment}"
}

# Create VPC only if it doesn't exist
resource "digitalocean_vpc" "app_network" {
  count = data.digitalocean_vpc.playground_vpc.id == "" ? 1 : 0
  
  name     = "playground-network-${var.environment}"
  region   = var.region
  ip_range = var.vpc_cidr
}

# Use the VPC ID from either the data source or the resource
locals {
  vpc_id = data.digitalocean_vpc.playground_vpc.id != "" ? data.digitalocean_vpc.playground_vpc.id : (length(digitalocean_vpc.app_network) > 0 ? digitalocean_vpc.app_network[0].id : "")
  git_sha = substr(var.git_sha, 0, 7)
}

# Load balancer droplet (created once, not part of CI/CD pipeline)
resource "digitalocean_droplet" "load_balancer" {
  # Only create if it doesn't exist or if explicitly recreated
  count    = var.recreate_load_balancer ? 1 : (data.digitalocean_droplet.existing_load_balancer.id == "" ? 1 : 0)
  
  name       = "playground-lb-${var.environment}"
  size       = var.lb_droplet_size
  image      = var.droplet_image
  region     = var.region
  ssh_keys   = [data.digitalocean_ssh_key.default.id]
  tags       = ["lb", var.environment]
  monitoring = var.monitoring
  vpc_uuid   = local.vpc_id
  
  # Ensure VPC is created first if needed
  depends_on = [digitalocean_vpc.app_network]
  
  connection {
    type        = "ssh"
    user        = "root"
    private_key = file(var.ssh_private_key_path)
    host        = self.ipv4_address
  }

  # This prevents this resource from being recreated during CI/CD
  lifecycle {
    prevent_destroy = true
    ignore_changes  = [tags, name]
  }
}

# Use a data source to check if load balancer exists
data "digitalocean_droplet" "existing_load_balancer" {
  name = "playground-lb-${var.environment}"
  # This will fail silently if the droplet doesn't exist
}
# Data source to find existing app server droplets
data "external" "existing_app_servers" {
  program = ["sh", "-c", "curl -s -X GET -H 'Content-Type: application/json' -H \"Authorization: Bearer ${var.do_token}\" \"https://api.digitalocean.com/v2/droplets?tag_name=app\" | jq -r '.droplets | map({id: .id, name: .name}) | .[] | select(.name | contains(\"playground-${var.environment}-${local.git_sha}\") | not) | {id: .id, name: .name}' "]
}

# Null resource to delete old app server droplets
resource "null_resource" "delete_old_app_servers" {
  depends_on = [digitalocean_droplet.app_server]

  provisioner "local-exec" {
    command = <<EOT
      echo "${data.external.existing_app_servers.result}" | jq -r '.[].id' | while read id; do
        curl -s -X DELETE \
          -H "Authorization: Bearer ${var.do_token}" \
          "https://api.digitalocean.com/v2/droplets/$id"
      done
    EOT
  }
}

# Application droplet (recreated on each deployment)
resource "digitalocean_droplet" "app_server" {
  name     = "playground-${var.environment}-${local.git_sha}"
  size     = var.droplet_size
  image    = var.droplet_image
  region   = var.region
  ssh_keys = [data.digitalocean_ssh_key.default.id]
  tags     = ["app", var.environment]
  monitoring = var.monitoring

  vpc_uuid = local.vpc_id
  
  # Ensure VPC is created first if needed
  depends_on = [digitalocean_vpc.app_network]

  connection {
    type        = "ssh"
    user        = "root"
    private_key = file(var.ssh_private_key_path)
    host        = self.ipv4_address
  }
}

data "digitalocean_ssh_key" "default" {
  name = var.ssh_key_name
}

# Firewall for load balancer - allows public HTTP/HTTPS access
resource "digitalocean_firewall" "lb_firewall" {
  name = "playground-lb-firewall-${var.environment}"

  droplet_ids = [data.digitalocean_droplet.existing_load_balancer.id != "" ? data.digitalocean_droplet.existing_load_balancer.id : (length(digitalocean_droplet.load_balancer) > 0 ? digitalocean_droplet.load_balancer[0].id : "")]

  # SSH
  inbound_rule {
    protocol         = "tcp"
    port_range       = "22"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  # HTTP
  inbound_rule {
    protocol         = "tcp"
    port_range       = "80"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  # HTTPS
  inbound_rule {
    protocol         = "tcp"
    port_range       = "443"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  # All outbound traffic
  outbound_rule {
    protocol              = "tcp"
    port_range            = "1-65535"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol              = "udp"
    port_range            = "1-65535"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol              = "icmp"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }
}

# Firewall for app server - only allows traffic from load balancer and SSH
resource "digitalocean_firewall" "app_firewall" {
  name = "playground-app-firewall-${var.environment}-${local.git_sha}"

  droplet_ids = [digitalocean_droplet.app_server.id]

  # SSH
  inbound_rule {
    protocol         = "tcp"
    port_range       = "22"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  # HTTP from load balancer only
  inbound_rule {
    protocol         = "tcp"
    port_range       = "8080"
    source_addresses = [data.digitalocean_droplet.existing_load_balancer.id != "" ? data.digitalocean_droplet.existing_load_balancer.ipv4_address : (length(digitalocean_droplet.load_balancer) > 0 ? digitalocean_droplet.load_balancer[0].ipv4_address : "") + "/32"]
  }

  # Frontend port from load balancer only
  inbound_rule {
    protocol         = "tcp"
    port_range       = "3000"
    source_addresses = [data.digitalocean_droplet.existing_load_balancer.id != "" ? data.digitalocean_droplet.existing_load_balancer.ipv4_address : (length(digitalocean_droplet.load_balancer) > 0 ? digitalocean_droplet.load_balancer[0].ipv4_address : "") + "/32"]
  }

  # All outbound traffic
  outbound_rule {
    protocol              = "tcp"
    port_range            = "1-65535"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol              = "udp"
    port_range            = "1-65535"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol              = "icmp"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  depends_on = [digitalocean_droplet.app_server]

  # Force recreation when git_sha changes
  lifecycle {
    replace_triggered_by = [var.git_sha]
  }
}

resource "local_file" "ansible_inventory" {
  content = templatefile("${path.module}/templates/inventory.tpl",
    {
      app_ip = digitalocean_droplet.app_server.ipv4_address
      lb_ip  = data.digitalocean_droplet.existing_load_balancer.id != "" ? data.digitalocean_droplet.existing_load_balancer.ipv4_address : (length(digitalocean_droplet.load_balancer) > 0 ? digitalocean_droplet.load_balancer[0].ipv4_address : "")
    }
  )
  filename = "../ansible/inventory.yml"
}

output "app_ip" {
  value = digitalocean_droplet.app_server.ipv4_address
  description = "The public IP address of the deployed app droplet"
}

output "load_balancer_ip" {
  value = data.digitalocean_droplet.existing_load_balancer.id != "" ? data.digitalocean_droplet.existing_load_balancer.ipv4_address : (length(digitalocean_droplet.load_balancer) > 0 ? digitalocean_droplet.load_balancer[0].ipv4_address : "")
  description = "The public IP address of the load balancer droplet (use this for Cloudflare DNS)"
}

output "app_url" {
  value = "https://${var.subdomain}.${var.domain_name}"
  description = "The URL to access the application"
}
