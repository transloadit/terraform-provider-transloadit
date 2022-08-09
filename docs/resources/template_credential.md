---
layout: "transloadit"
page_title: "Transloadit: template credential"
description: |-
Manages Transloadit template credential
---

# Resource: transloadit_template_credential

Manages Transloadit Templates Credential
For additional details please refer to the [Transloadit documentation](https://transloadit.com/docs/topics/template-credentials/)

## Example Usage

```hcl
provider "transloadit" {

}

resource "transloadit_template_credential" "s3" {
  name = "S3_CREDENTIALS"
  type = "s3"
  content = <<EOT
  {
	"bucket" : "mybucket",
	"bucket_region" : "us-east-1",
    "key" : "mykey",
	"secret" : "mysecret"
   }
EOT
}

resource "transloadit_template" "my-terraform-template" {
	name     = "my-terraform-template"
	template = <<EOT
{
  "steps": {
    ":original": {
      "robot": "/upload/handle"
    },
    [... different steps ...]
    "exported": {
      "use": [ ":original", "encoded", "thumbed"], 
      "credentials": "${transloadit_template_credential.s3.name}",
      "robot": "/s3/store"
    }
  }
}
EOT
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) name of the Template credential to be added
- `type` - (Required) Type of Template credential to be added
- `content` - (Required) Json of the template credential . Go to the [documentation](https://transloadit.com/docs/transcoding/) to have informations about the parameters needed.
  Example : 
  - s3 credentials : https://transloadit.com/docs/transcoding/file-importing/s3-import/#credentials
  - backblaze : https://transloadit.com/docs/transcoding/file-importing/backblaze-import/#credentials
  - sftp : https://transloadit.com/docs/transcoding/file-importing/sftp-import/#credentials

## Attributes Reference

- `id` - The Template ID reference from Transloadit, e.g. `908ab54f2c4b479184db637709320c85`
- `name` - The Template Slug reference from Transloadit, e.g. `my-terraform-template`
- `type` -  Json of the template credential
- `content` - The Template's Assembly Instructions as a JSON string, e.g. `{ "steps": { ... } }`
- `created` - Creation date of the templace 

## Update 

Currently update is not supported but will probably be available in a future release

## Import

Template credential can be imported using the `id`, e.g.

```bash
$ terraform import transloadit_template_credential.my_template 908ab54f2c4b479184db637709320c85
```

