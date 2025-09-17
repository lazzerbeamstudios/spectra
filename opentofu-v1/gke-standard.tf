resource "google_container_cluster" "gke_standard" {
  name = "${var.gke_name}-standard"

  network  = var.vpc_name
  location = var.zone

  deletion_protection = false

  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "node_pool" {
  name = var.node_pool_name_1

  cluster  = google_container_cluster.gke_standard.name
  location = var.zone

  node_count = var.node_pool_node_count_1

  node_config {
    machine_type = var.node_pool_machine_type_1
    disk_size_gb = var.node_pool_disk_size_gb_1
    preemptible  = false
  }
}
