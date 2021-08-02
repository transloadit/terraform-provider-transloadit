terraform {
  required_version = ">= 1.0.0"
  required_providers {
    transloadit = {
      source  = "transloadit/transloadit"
      version = "0.4.0"
    }
  }
}

variable "tl_urltransform_key" { type = string }
variable "tl_urltransform_secret" { type = string }
variable "tl_preprocess_key" { type = string }
variable "tl_preprocess_secret" { type = string }

provider "transloadit" {
  auth_key    = var.tl_urltransform_key
  auth_secret = var.tl_urltransform_secret
  alias       = "urltransform"
}
provider "transloadit" {
  auth_key    = var.tl_preprocess_key
  auth_secret = var.tl_preprocess_secret
  alias       = "preprocess"
}


module "urltransform" {
  source    = "./urltransform"
  providers = {
    transloadit = transloadit.urltransform
  }
}

module "preprocess" {
  source    = "./preprocess"
  providers = {
    transloadit = transloadit.preprocess
  }
}

// Steps:
// - Create `./urltransform/main.tf`, with at least a `terraform.required_providers.transloadit` block
// - Run `terraform init`
// - Move your template resources to `./urltransform/main.tf`
// - Run `terraform state mv 'transloadit_template.resize-img' 'module.urltransform.transloadit_template.resize-img'` (or any other Template)

output "cdn-resize-id" { value = module.urltransform.resize-img-id }