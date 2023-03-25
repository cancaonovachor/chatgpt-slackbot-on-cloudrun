resource "google_cloud_run_service" "default" {
  name     = "slack-chatgpt-bot"
  location = var.region

  template {
    spec {
      containers {
        image = "asia-northeast1-docker.pkg.dev/${var.project_id}/slack-chatgpt-bot"

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
  depends_on = [
    google_cloudbuild_trigger.slack_chatgpt_bot
  ]
}


