terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "2.6.3"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}