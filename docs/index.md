---
layout: "transloadit"
page_title: "Provider: Transloadit"
description: |-
  The Transloadit provider is used to manage Transloadit resources. The provider needs to be configured with the proper credentials before it can be used.
---

# Transloadit Provider

The Transloadit provider is used to manage Transloadit Templates, by which you can declaratively configure automated media encoding pipelines.
The provider needs to be configured with the proper credentials before it can be used.

## Examples

Follow these steps to get started:

- After signing up, Get your [Transloadit credentials](https://transloadit.com/c/<MYACCOUNT>/template-credentials/)
- Create a new directory, cd into it, create a file called `main.tf`
- Initialize a Terraform working directory: `terraform init`
- Generate and show the execution plan: `terraform plan`
- Build the infrastructure: `terraform apply`

```hcl
provider "transloadit" {
  auth_key    = "<TRANSLOADIT-AUTH-KEY>"
  auth_secret = "<TRANSLOADIT-AUTH-SECRET>"
  version     = "0.2.0"
}

resource "transloadit_template" "to-ipad-with-thumbnails" {
  name     = "to-ipad-with-thumbnails"
  template = <<EOT
{
  "steps": {
    ":original": {
      "robot": "/upload/handle"
    },
    "encoded": {
      "use": ":original",
      "robot": "/video/encode",
      "preset": "ipad-high"
    },
    "thumbed": {
      "use": ":original",
      "robot": "/video/thumbs",
      "count": 8
    }
  }
}
EOT
}
```

## Authentication

Transloadit authentication is based on an **access key** and a **secret key**.

The Transloadit provider offers two ways of providing these credentials. The following methods are checked, in this order:

1. [Static credentials](#static-credentials)
2. [Environment variables](#environment-variables)

### Static credentials

!> **Warning** Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this file, or the state file, ever be committed to a public version control system.

Static credentials can be provided by adding `access_key` and `secret_key` attributes in-line in the Transloadit provider block:

Example:

```hcl
provider "transloadit" {
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}
```

### Environment variables

You can provide your credentials via the `TRANSLOADIT_AUTH_KEY`, `TRANSLOADIT_AUTH_SECRET` environment variables.

Example:

```hcl
provider "transloadit" {}
```

Usage:

```bash
export TRANSLOADIT_AUTH_KEY="my-access-key"
export TRANSLOADIT_AUTH_SECRET="my-secret-key"
terraform apply
```

## Arguments Reference

In addition to [generic provider arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g. `alias` and `version`), the following arguments are supported in the Transloadit provider block:

- `auth_key` - (Optional) The Transloadit access key. It must be provided, but it can also be sourced from
the `TRANSLOADIT_AUTH_KEY` [environment variable](#environment-variables)

- `auth_secret` - (Optional) The Transloadit secret key. It must be provided, but it can also be sourced from
the `TRANSLOADIT_AUTH_SECRET` [environment variable](#environment-variables)

- `endpoint` - (Optional) The Transloadit API endpoint. This defaults to `https://api2.transloadit.com` and automatically georoutes you to the fastest API location. You shouldn't have to change this unless you only deal with a single availability region or are a plugin contributor targeting local installs.
