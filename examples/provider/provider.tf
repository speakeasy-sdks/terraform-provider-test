terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "1.7.0"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}