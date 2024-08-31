resource "google_pubsub_topic" "user_added" {
  name = var.user_added_topic_name

  labels = {
    subject = "user"
  }

  message_retention_duration = "86600s"
}

resource "google_pubsub_subscription" "user_added" {
  name  = var.user_added_subscription_name
  topic = google_pubsub_topic.user_added.id

  labels = {
    subject = "user"
  }

  message_retention_duration = "1200s"
  retain_acked_messages      = true

  ack_deadline_seconds = 20

  expiration_policy {
    ttl = "300000.5s"
  }

  retry_policy {
    minimum_backoff = "10s"
  }

  enable_message_ordering    = false
}

resource "google_cloud_run_v2_service" "books_app" {
  name     = var.cloudrun_service_name
  location = "us-central1"
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

resource "google_artifact_registry_repository" "books_app_registry" {
  location = "us-central1"
  repository_id = "books-app-registry"
  format = "DOCKER"
}

resource "google_sql_database" "user_database" {
  name     = var.books_database_name
  instance = google_sql_database_instance.user_database_instance.name
}

resource "google_sql_database_instance" "user_database_instance" {
  name             = var.books_database_instance_name
  database_version = "POSTGRES_14"
  deletion_protection = false

  settings {
    tier = "db-f1-micro"
  }
}

resource "google_firestore_database" "firestore_database" {
  name        = var.firestore_database_name
  location_id = "us-central1"
  type        = "FIRESTORE_NATIVE"
}
