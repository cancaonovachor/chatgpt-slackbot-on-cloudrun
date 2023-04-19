terraform {
  required_version = "~> 1.4.0"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.56.0"
    }
  }
}


provider "google" {
  region  = var.region
  project = var.project_id
}
