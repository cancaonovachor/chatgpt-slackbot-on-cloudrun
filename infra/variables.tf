variable "project_id" {
  type        = string
  description = "The project ID to deploy the resources."
}

variable "region" {
  type        = string
  description = "The region for deploying resources."
  default     = "asia-northeast1"
}

variable "openai_api_key" {
  type        = string
  description = "The OpenAI API key."
  sensitive   = true
}

variable "slack_signing_secret" {
  type        = string
  description = "The Slack signing secret."
  sensitive   = true
}
