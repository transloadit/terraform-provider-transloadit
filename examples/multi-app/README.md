## Provisioning Multiple Transloadit Apps with Terraform

If want to manage mulitple Transloadit Apps with Terraform, this folder contains an example for how to structure that with Terraform
Modules.

### Migrating from Single App

If you were previously managing one app and not using Modules, here are the basic
steps to migrate without destroying Templates along the way. 

These steps will have to be modified to match your environment, and of course
always carefully plan and make sure nothing destructive happens, before applying.

We assume in this scenario that you just had one `main.tf`, or maybe one `transloadit.tf` controlling the `urltransform` App, and you are now adding a `preprocessing` App.

Steps:

1. First make sure your local Terraform recipes are up to date with production by running `terraform plan`. Fix any drift if you encounter it.
1. Create a Module `./urltransform/main.tf`, with at least a `terraform.required_providers.transloadit` block (see example file)
1. Create a Module `./preprocessing/main.tf`, with at least a `terraform.required_providers.transloadit` block (see example file)
1. Add one Transloadit Provider Alias for each App in `./main.tf`, and invoke both Modules, passing in their appropriate App credentials, which are `variable`s, sourced via `env.sh`
1. Run `terraform init` to initialize the Modules
1. Move your exisitng `urltransform` Template resources to `./urltransform/main.tf`
1. Run `terraform state mv 'transloadit_template.resize-img' 'module.urltransform.transloadit_template.resize-img'` (or any other Template instead of `resize-img`)
1. After all resources have succesfully been moved into the App, run: `terraform plan` and `terraform apply`.
1. There should be no changes.
1. Assuming you're adding a brand new App, you can now start adding any new Template into the new `preprocessing` App.
1. You can `ouput` any resource attribute from the Module, back to the root `main.tf`, which in turn can `output` the Module attribute now containing the information (such as the Template ID created).
