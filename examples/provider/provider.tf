terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "1.3.3"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}