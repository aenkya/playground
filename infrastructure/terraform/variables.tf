variable "do_token" {
  description = "DigitalOcean API token"
  type        = string
  sensitive   = true
}

variable "environment" {
  description = "Environment (prod/staging)"
  type        = string
  default     = "prod"
}

variable "region" {
  description = "DigitalOcean region"
  type        = string
  default     = "nyc1"
}

variable "droplet_size" {
  description = "Droplet size"
  type        = string
  default     = "s-1vcpu-1gb"
}

variable "droplet_image" {
  description = "Droplet base image"
  type        = string
  default     = "ubuntu-24-04-x64"
}

variable "vpc_cidr" {
  description = "VPC CIDR range"
  type        = string
  default     = "10.10.10.0/24"
}

variable "ssh_key_name" {
  description = "Name of SSH key in DigitalOcean"
  type        = string
}

variable "ssh_private_key_path" {
  description = "Path to SSH private key"
  type        = string
}

variable "monitoring" {
  description = "Enable monitoring"
  type        = bool
  default     = false
}