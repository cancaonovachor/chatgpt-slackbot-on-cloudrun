resource "google_storage_bucket" "cloudbuild" {
  name          = "${var.project_id}-cloudbuild"
  location      = var.region
  force_destroy = true
}
