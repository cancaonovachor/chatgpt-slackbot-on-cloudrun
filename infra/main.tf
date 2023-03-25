terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

locals {
  image_name = "gcr.io/${var.project_id}/slack-chatbot"
}

module "cloudrun" {
  source  = "terraform-google-modules/cloud-run/google"
  version = "~> 2.0"

  project_id            = var.project_id
  name                  = "slack-chatbot"
  location              = var.region
  image                 = local.image_name
  allow_unauthenticated = true

  env_vars = {
    OPENAI_API_KEY       = var.openai_api_key
    SLACK_SIGNING_SECRET = var.slack_signing_secret
  }
}

output "cloudrun_url" {
  value = module.cloudrun.service_url
}
