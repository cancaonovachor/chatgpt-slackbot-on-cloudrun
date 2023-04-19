resource "google_pubsub_topic" "chatgpt_bot_topic" {
  name = "chatgpt-bot-topic"
}

resource "google_pubsub_subscription" "chatgpt_bot_subscription" {
  name                 = "chatgpt-bot-subscription"
  topic                = google_pubsub_topic.chatgpt_bot_topic.name
  ack_deadline_seconds = 20
  push_config {
    push_endpoint = google_cloud_run_service.gptapp_cloud_run_service.status[0].url
  }
  depends_on = [
    google_pubsub_topic.chatgpt_bot_topic,
    google_cloud_run_service.gptapp_cloud_run_service
  ]
}
