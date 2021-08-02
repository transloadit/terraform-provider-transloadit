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

output "resize-img-id" { value = transloadit_template.resize-img.id }