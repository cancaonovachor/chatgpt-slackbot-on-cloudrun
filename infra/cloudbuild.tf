resource "google_cloudbuild_trigger" "slack_chatgpt_bot" {
  provider = google

  name     = "slack-chatgpt-bot"
  disabled = false

  included_files = [
    "Dockerfile",
    "app.py",
    "cloudbuild.yaml",
    "requirements.txt",
  ]
  github {
    owner = "cancaonovachor"
    name  = "chatgpt-slackbot-on-cloudrun"
    push {
      branch = "^main$"
    }
  }

  filename = "cloudbuild.yaml"

  substitutions = {
    _REGION = var.region
  }
}
