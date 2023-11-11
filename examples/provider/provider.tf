terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "2.3.1"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}