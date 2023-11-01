terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "1.12.0"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}