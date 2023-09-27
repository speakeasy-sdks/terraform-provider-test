terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "1.8.1"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}