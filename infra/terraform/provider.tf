terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "5.42.0"
    }
  }
}

provider "google" {
  project = "global-impulse-433212-j3"
  region = "us-central1"
  zone = "us-central1-a"
  credentials = "key.json"
}

