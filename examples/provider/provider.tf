terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "4.0.0"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}