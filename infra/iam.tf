resource "google_cloud_run_service_iam_policy" "default" {
  location = google_cloud_run_service.default.location
  project  = google_cloud_run_service.default.project
  service  = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.default.policy_data
}

data "google_iam_policy" "default" {
  binding {
    role = "roles/run.invoker"

    members = [
      "allUsers",
    ]
  }
}
