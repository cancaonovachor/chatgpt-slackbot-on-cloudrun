output "pubsubapp_cloudrun_url" {
  value = google_cloud_run_service.pubsubapp_cloud_run_service.status[0].url
}
