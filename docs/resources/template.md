---
layout: "transloadit"
page_title: "Transloadit: template"
description: |-
  Manages Transloadit template
---

# Resource: transloadit_template

Manages Transloadit Templates. 
For additional details please refer to the [Transloadit documentation](https://transloadit.com/docs/)

## Example Usage

```hcl
provider "transloadit" {

}

resource "transloadit_template" "my-terraform-template" {
	name     = "my-terraform-template"
	template = <<EOT
{
  "steps": {
    ":original": {
      "robot": "/upload/handle"
    },
    "encoded": {
      "use": ":original",
      "robot": "/video/encode",
      "preset": "ipad-high",
      "ffmpeg_stack": "v3.3.3"
    },
    "thumbed": {
      "use": ":original",
      "robot": "/video/thumbs",
      "count": 4,
      "ffmpeg_stack": "v3.3.3"
    },
    "exported": {
      "use": [ ":original", "encoded", "thumbed"], 
      "credentials": "YOUR_S3_CREDENTIALS",
      "robot": "/s3/store"
    }
  }
}
EOT
}
```

~> If you use [Assembly Variables](https://transloadit.com/docs/#assembly-variables) such as `${fields.width}`, make sure to [escape them](https://www.terraform.io/docs/configuration-0-11/interpolation.html) via double dollar signs: `$${fields.width}`.

## Argument Reference

The following arguments are supported:

- `name` - (Required) name of the Template to be added
- `template` - (Required) JSON string of the Template's Assembly Instructions. The top-level object must be the JSON object with a `"steps"` key. For more details, see [Assembly Instructions](https://transloadit.com/docs/#assembly-instructions)
- `require_signature_auth` - (Optional - boolean) Use true to deny requests that do not include a signature. With [Signature Authentication](https://transloadit.com/docs/#signature-authentication) you can ensure no one else is sending requests on your behalf. Default value: false

## Attributes Reference 

- `id` - The Template ID reference from Transloadit, e.g. `908ab54f2c4b479184db637709320c85`
- `name` - The Template Slug reference from Transloadit, e.g. `my-terraform-template`
- `template` - The Template's Assembly Instructions as a JSON string, e.g. `{ "steps": { ... } }`

## Import

Templates can be imported using the `id`, e.g.

```bash
$ terraform import transloadit_template.my_template 908ab54f2c4b479184db637709320c85
```

