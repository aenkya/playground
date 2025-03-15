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

# Load balancer droplet (created once, not part of CI/CD pipeline)
resource "digitalocean_droplet" "load_balancer" {
  name       = "playground-lb-${var.environment}"
  size       = var.lb_droplet_size
  image      = var.droplet_image
  region     = var.region
  ssh_keys   = [data.digitalocean_ssh_key.default.id]
  tags       = ["lb", var.environment]
  monitoring = var.monitoring
  vpc_uuid   = digitalocean_vpc.app_network.id

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

# Application droplet (recreated on each deployment)
resource "digitalocean_droplet" "app_server" {
  name     = "playground-${var.environment}"
  size     = var.droplet_size
  image    = var.droplet_image
  region   = var.region
  ssh_keys = [data.digitalocean_ssh_key.default.id]
  tags     = ["app", var.environment]
  monitoring = var.monitoring

  vpc_uuid = digitalocean_vpc.app_network.id

  connection {
    type        = "ssh"
    user        = "root"
    private_key = file(var.ssh_private_key_path)
    host        = self.ipv4_address
  }
}

resource "digitalocean_vpc" "app_network" {
  name     = "playground-network-${var.environment}"
  region   = var.region
  ip_range = var.vpc_cidr
}

data "digitalocean_ssh_key" "default" {
  name = var.ssh_key_name
}

# Firewall for load balancer - allows public HTTP/HTTPS access
resource "digitalocean_firewall" "lb_firewall" {
  name = "playground-lb-firewall-${var.environment}"

  droplet_ids = [digitalocean_droplet.load_balancer.id]

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
  name = "playground-app-firewall-${var.environment}"

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
    source_addresses = ["${digitalocean_droplet.load_balancer.ipv4_address}/32"]
  }

  # Frontend port from load balancer only
  inbound_rule {
    protocol         = "tcp"
    port_range       = "3000"
    source_addresses = ["${digitalocean_droplet.load_balancer.ipv4_address}/32"]
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

resource "local_file" "ansible_inventory" {
  content = templatefile("${path.module}/templates/inventory.tpl",
    {
      app_ip = digitalocean_droplet.app_server.ipv4_address
      lb_ip  = digitalocean_droplet.load_balancer.ipv4_address
    }
  )
  filename = "../ansible/inventory.yml"
}

output "droplet_ip" {
  value = digitalocean_droplet.app_server.ipv4_address
  description = "The public IP address of the deployed app droplet"
}

output "load_balancer_ip" {
  value = digitalocean_droplet.load_balancer.ipv4_address
  description = "The public IP address of the load balancer droplet (use this for Cloudflare DNS)"
}

output "app_url" {
  value = "https://${var.subdomain}.${var.domain_name}"
  description = "The URL to access the application"
}
