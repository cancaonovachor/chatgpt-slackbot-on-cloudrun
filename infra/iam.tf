resource "google_cloud_run_service_iam_policy" "gptapp_cloud_run_service" {
  location = google_cloud_run_service.gptapp_cloud_run_service.location
  project  = google_cloud_run_service.gptapp_cloud_run_service.project
  service  = google_cloud_run_service.gptapp_cloud_run_service.name

  policy_data = data.google_iam_policy.gptapp_cloud_run_service.policy_data
}

data "google_iam_policy" "gptapp_cloud_run_service" {
  binding {
    role = "roles/run.invoker"

    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "pubsubapp_cloud_run_service" {
  location = google_cloud_run_service.pubsubapp_cloud_run_service.location
  project  = google_cloud_run_service.pubsubapp_cloud_run_service.project
  service  = google_cloud_run_service.pubsubapp_cloud_run_service.name

  policy_data = data.google_iam_policy.pubsubapp_cloud_run_service.policy_data
}

data "google_iam_policy" "pubsubapp_cloud_run_service" {
  binding {
    role = "roles/run.invoker"

    members = [
      "allUsers",
    ]
  }
}
