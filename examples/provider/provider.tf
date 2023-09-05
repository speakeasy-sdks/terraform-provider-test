terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "1.6.2"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}