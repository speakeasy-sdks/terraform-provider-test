terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "1.10.0"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}