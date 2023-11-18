terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "2.4.0"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}