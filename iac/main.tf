resource "google_service_account" "default" {
  account_id   = var.SERVICE_ACCOUNT
  display_name = "Quizzer Service Account"
}

resource "google_container_cluster" "primary" {
  name     = "quizzer"
  location = var.region

  remove_default_node_pool = true
  initial_node_count       = 2
}

resource "google_container_node_pool" "primary_nodes" {
  name       = "pool-1"
  location   = var.region
  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "e2-medium"

    service_account = google_service_account.default.email
    oauth_scopes    = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
