resource "google_cloudbuild_trigger" "slack-bot" {
  provider = google

  name     = "slack-bot"
  disabled = false

  included_files = [
    "Dockerfile",
    "app.py",
    "cloudbuild.yaml",
    "requirements.txt",
  ]

  trigger_template {
    branch_name = "main"
    repo_name   = "https://github.com/cancaonovachor/chatgpt-slackbot-on-cloudrun"
  }

  filename = "cloudbuild.yaml"

  substitutions = {
    _REGION = var.region
  }
}
