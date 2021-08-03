terraform {
  required_version = ">= 1.0.0"
  required_providers {
    transloadit = {
      source  = "transloadit/transloadit"
      version = "0.4.0"
    }
  }
}

// Define our `urltransform` App and its credentials
variable "tl_urltransform_key"    { type = string }
variable "tl_urltransform_secret" { type = string }
provider "transloadit" {
  auth_key    = var.tl_urltransform_key
  auth_secret = var.tl_urltransform_secret
  alias       = "urltransform"
}

// Define our `preprocess` App and its credentials
variable "tl_preprocess_key"    { type = string }
variable "tl_preprocess_secret" { type = string }
provider "transloadit" {
  auth_key    = var.tl_preprocess_key
  auth_secret = var.tl_preprocess_secret
  alias       = "preprocess"
}

resource "transloadit_template" "resize-img" {
  provider = transloadit.urltransform // <-- In which App to save the template
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

resource "transloadit_template" "video-encode" {
  provider = transloadit.preprocess // <-- In which App to save the template
	name     = "video-encode"
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
    }
  }
}
EOT
}

output "cdn-resize-id" { value = transloadit_template.resize-img.id }