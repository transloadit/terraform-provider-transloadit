---
layout: "transloadit"
page_title: "Transloadit: template"
description: |-
  Manages Transloadit template
---

# transloadit_template

Manages Transloadit templates
For additional details please refer to [API documentation](https://transloadit.com/docs/)

## Example Usage

```hcl
provider "transloadit" {

}

resource "transloadit_template" "test" {
	name = "templatebasic1"
	template = <<EOT
    {
  "steps": {
    ":original": {
      "robot": "/upload/handle"
    },
    "encoded": {
      "preset": "ipad-high",
      "robot": "/video/encode",
      "use": ":original",
      "ffmpeg_stack": "v3.3.3"
    },
    "exported": {
      "credentials": "YOUR_S3_CREDENTIALS",
      "robot": "/s3/store",
      "use": [
        "encoded",
        "thumbed"
      ]
    },
    "thumbed": {
      "count": 4,
      "robot": "/video/thumbs",
      "use": ":original",
      "ffmpeg_stack": "v3.3.3"
    }
  }
}
EOT
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) name of the template to be added
- `template` - (Required) Json String of the template content. The top-level object must be the Json Object with a "steps" key . For more details, see [Assembly instructions](https://transloadit.com/docs/#assembly-instructions)

## Import

Templates can be imported using the `id`, e.g.

```
$ terraform import transloadit_template.my_template 908ab54f2c4b479184db637709320c85
```

