resource "google_cloud_run_service" "default" {
  name     = "slack-chatbot"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/slack-chatbot"

        env {
          name  = "OPENAI_API_KEY"
          value = var.openai_api_key
        }

        env {
          name  = "SLACK_SIGNING_SECRET"
          value = var.slack_signing_secret
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}


