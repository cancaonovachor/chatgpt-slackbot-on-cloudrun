resource "google_cloud_run_service" "gptapp_cloud_run_service" {
  name     = local.service_name_gptapp
  location = var.region

  template {
    spec {
      containers {
        image = "asia-northeast1-docker.pkg.dev/${var.project_id}/${local.service_name}/${local.service_name_gptapp}/image:${null_resource.gptapp_image_build.id}"
        env {
          name  = "OPENAI_API_KEY"
          value = var.openai_api_key
        }
        env {
          name  = "SLACK_BOT_TOKEN"
          value = var.slack_token
        }
        env {
          name  = "PROJECT_ID"
          value = var.project_id
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
  depends_on = [
    null_resource.gptapp_image_build
  ]
}


resource "google_cloud_run_service" "pubsubapp_cloud_run_service" {
  name     = local.service_name_pubsubapp
  location = var.region

  template {
    spec {
      containers {
        image = "asia-northeast1-docker.pkg.dev/${var.project_id}/${local.service_name}/${local.service_name_pubsubapp}/image:${null_resource.pubsubapp_image_build.id}"
        env {
          name  = "PUBSUB_TOPIC"
          value = google_pubsub_topic.chatgpt_bot_topic.name
        }
        env {
          name  = "PROJECT_ID"
          value = var.project_id
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
  depends_on = [
    null_resource.pubsubapp_image_build,
    google_pubsub_topic.chatgpt_bot_topic
  ]
}




