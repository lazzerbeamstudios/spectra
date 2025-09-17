# project
variable "project" {
  default = "[project]-gcp"
}

# region
variable "region" {
  default = "us-central1"
}

# zone
variable "zone" {
  default = "us-central1-a"
}

# vpc
variable "vpc_name" {
  default = "vpc-v1"
}

# gke
variable "gke_name" {
  default = "gke-v1"
}

# node pool 1
variable "node_pool_name_1" {
  default = "node-pool-1"
}

variable "node_pool_node_count_1" {
  default = 2
}

variable "node_pool_machine_type_1" {
  default = "e2-small"
}

variable "node_pool_disk_size_gb_1" {
  default = 25
}
