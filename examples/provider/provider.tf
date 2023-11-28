terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "2.4.1"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}