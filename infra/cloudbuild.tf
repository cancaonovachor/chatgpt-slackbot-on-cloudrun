# gpt用コードをlocalでbuild
resource "null_resource" "gptapp_image_build" {
  # 再デプロイをしたいとき用
  triggers = {
    trigger = "1"
  }
  provisioner "local-exec" {
    working_dir = "../application/gpt-app"
    interpreter = ["bash", "-c"]
    command = join(" ", [
      "gcloud builds submit",
      "--config cloudbuild.yaml .",
    ])
    on_failure = fail
  }
  depends_on = [
    google_project_service.apis,
    google_artifact_registry_repository.slack_chatgpt_bot
  ]
}

# pubsubレシーブ用コードをlocalでbuild
resource "null_resource" "pubsubapp_image_build" {
  # 再デプロイをしたいとき用
  triggers = {
    trigger = "1"
  }
  provisioner "local-exec" {
    working_dir = "../application/pubsub-app"
    interpreter = ["bash", "-c"]
    command = join(" ", [
      "gcloud builds submit",
      "--config cloudbuild.yaml .",
    ])
    on_failure = fail
  }
  depends_on = [
    google_project_service.apis,
    google_artifact_registry_repository.slack_chatgpt_bot
  ]
}

