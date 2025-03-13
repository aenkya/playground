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

resource "digitalocean_droplet" "app_server" {
  name     = "playground-${var.environment}"
  size     = var.droplet_size
  image    = var.droplet_image
  region   = var.region
  ssh_keys = [data.digitalocean_ssh_key.default.id]
  tags     = ["app", var.environment]

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

resource "local_file" "ansible_inventory" {
  content = templatefile("${path.module}/templates/inventory.tpl",
    {
      app_ip = digitalocean_droplet.app_server.ipv4_address
    }
  )
  filename = "../ansible/inventory.yml"
}
