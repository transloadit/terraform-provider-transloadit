terraform {
  required_providers {
    transloadit = {
      source  = "transloadit/transloadit"
      version = "0.4.0"
    }
  }
}

resource "transloadit_template" "resize-img" {
	name     = "resize-img"
	template = <<EOT
{
  "steps": {
    "imported": {
      "robot": "/s3/import",
      "credentials": "my_s3_$${assembly.region}",
      "path": "/onthefly/$${fields.input}"
    },
    "resized": {
      "use": "imported",
      "robot": "/image/resize",
      "width": "$${fields.w}",
      "imagemagick_stack": "v2.0.7"
    },
    "served": {
      "use": "resized",
      "robot": "/file/serve"
    }
  }
}
EOT
}

output "resize-img-id" { value = transloadit_template.resize-img.id }