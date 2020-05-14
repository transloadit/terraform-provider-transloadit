---
layout: "transloadit"
page_title: "Provider: Transloadit"
description: |-
  The Transloadit provider is used to manage Transloadit resources. The provider needs to be configured with the proper credentials before it can be used.
---

# Transloadit Provider

The Transloadit provider is used to manage Transloadit resources.
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Examples

Here is an example that will setup a web server with an additional volume, a public IP and a security group.

You can test this config by creating a `test.tf` and run terraform commands from this directory:

- Get your [Transloadit credentials](https://transloadit.com/c/<MYACCOUNT>/template-credentials)
- Initialize a Terraform working directory: `terraform init`
- Generate and show the execution plan: `terraform plan`
- Build the infrastructure: `terraform apply`

```hcl
provider "transloadit" {
  auth_key    = "<TRANSLOADIT-AUTH-KEY>"
  auth_secret = "<TRANSLOADIT-AUTH-SECRET>"
  version     = "0.1.0"
}
```

## Authentication

The Transloadit authentication is based on an **access key** and a **secret key**.

The Transloadit provider offers two ways of providing these credentials. The following methods are supported, in this priority order:

1. [Static credentials](#static-credentials)
2. [Environment variables](#environment-variables)

### Static credentials

!> **Warning**: Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this file, or the state file, ever be committed to a public version control system.

Static credentials can be provided by adding `access_key` and `secret_key` attributes in-line in the Transloadit provider block:

Example:

```hcl
provider "transloadit" {
  access_key = "my-access-key"
  secret_key = "my-secret-key"
  version    = "0.1.0"
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
$ export TRANSLOADIT_AUTH_KEY="my-access-key"
$ export TRANSLOADIT_AUTH_SECRET="my-secret-key"
$ terraform plan
```

## Arguments Reference

In addition to [generic provider arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g. `alias` and `version`), the following arguments are supported in the Transloadit provider block:

- `auth_key` - (Optional) The Transloadit access key. It must be provided, but it can also be sourced from
the `TRANSLOADIT_AUTH_KEY` [environment variable](#environment-variables)

- `auth_secret` - (Optional) The Transloadit secret key. It must be provided, but it can also be sourced from
the `TRANSLOADIT_AUTH_SECRET` [environment variable](#environment-variables)

- `endpoint` - (Optional) The Transloadit API endpoint. For standard usages, this argument don't have to be used
