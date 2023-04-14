terraform {
  cloud {
    organization = "ProtoCloud"

    workspaces {
      name = "quizzer"
    }
  }

  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.61.0"
    }

    kubernetes = {
      source = "hashicorp/kubernetes"
      version = "2.19.0"
    }
  }

  required_version = ">= 1.1.9"
}


provider "google" {
  project = var.PROJECT_ID
}
