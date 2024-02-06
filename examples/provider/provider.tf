terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "3.0.1"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}