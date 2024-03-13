terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "4.1.0"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}