data "google_service_account" "default" {
  account_id   = var.SERVICE_ACCOUNT
}

resource "google_container_cluster" "primary" {
  name     = "quizzer"
  location = var.location

  remove_default_node_pool = true
  initial_node_count       = 3
}

resource "google_container_node_pool" "primary_nodes" {
  name       = "pool-1"
  location   = var.location
  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "e2-medium"

    service_account = data.google_service_account.default.email
    oauth_scopes    = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
