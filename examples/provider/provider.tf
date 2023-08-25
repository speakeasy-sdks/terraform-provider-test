terraform {
  required_providers {
    AcmeTerraform = {
      source  = "vitor-test/AcmeTerraform"
      version = "1.3.4"
    }
  }
}

provider "AcmeTerraform" {
  # Configuration options
}