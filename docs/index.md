---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "repo Provider"
subcategory: ""
description: |-
  User API for Speakeasy template service: The Rest Template API is an API used for instrucitonal purposes.
---

# repo Provider

User API for Speakeasy template service: The Rest Template API is an API used for instrucitonal purposes.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `server_url` (String) Server URL (defaults to http://localhost:8080)