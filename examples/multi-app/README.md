## Provisioning Multiple Transloadit Apps with Terraform

If want to manage mulitple Transloadit Apps with Terraform, this folder contains an example for how to structure that with Terraform
Modules.

We assume in this scenario that you just had one `main.tf`, or maybe one `transloadit.tf` controlling the `urltransform` App, and you are now adding a `preprocessing` App.

You were implicitly using one App's Auth Key and Secret, without naming the App in your original HCL.

In the new situation, we'll specifically give each provider the name of the App, using the [`alias`](https://www.terraform.io/docs/language/providers/configuration.html#alias-multiple-provider-configurations) argument.

This means giving an `alias` to the existing Provider, and adding a new one.

We'll pass in credentials via `TF_VAR_`. They are set from `env.sh`, but credentials can be passed via any other Terraform supported mechanism.

In the Template resources (`transloadit_template`), we explicitly set which App they should be saved in, using the `provider` argument.

For the full code check out [`main.tf`](./main.tf), and [`env.sh`](./env.sh).