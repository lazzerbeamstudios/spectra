resource "google_container_cluster" "gke_autopilot" {
  name = "${var.gke_name}-autopilot"

  network  = google_compute_network.vpc.name
  location = var.region

  deletion_protection = false
  enable_autopilot    = true
}
