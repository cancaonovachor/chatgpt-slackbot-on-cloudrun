# Artifact Registryのリポジトリ
resource "google_artifact_registry_repository" "slack_chatgpt_bot" {
  location      = "asia-northeast1"
  repository_id = local.service_name
  description   = "chatgpt slack bot server"
  format        = "DOCKER"
  depends_on = [
    google_project_service.apis,
  ]
}
