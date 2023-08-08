terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "1.1.0"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}