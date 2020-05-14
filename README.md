# Transloadit Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x

## Installing the provider

The recommended way to install *terraform-provider-transloadit* is to use the binary
distributions from the [Releases](https://github.com/transloadit/terraform-provider-transloadit/releases) page. The packages are available for Linux, macOS (darwin) and Windows (all on amd64 architecture).

Download and uncompress the latest release for your OS. This example uses the Linux binary.

```bash
mkdir -p ~/.terraform.d/plugins/
cd ~/.terraform.d/plugins/
curl -sSL https://github.com/transloadit/terraform-provider-transloadit/releases/download/v0.1.0/terraform-provider-transloadit_linux_amd64.tar.gz |tar xvz
```

After saving and exctracting the provider to your plugins directory, you'll want to initialize Terraform in your project so it can find the plugin.

```bash
terraform init
```

On windows, you have to store the provider in `%APPDATA%\terraform.d\plugins` 

## Using the provider

Here's a quick example. More detailed instructions can be found in the [website directory](./website/).

In `main.tf`:

```hcl
provider "transloadit" {
  auth_key    = "<TRANSLOADIT-AUTH-KEY>"
  auth_secret = "<TRANSLOADIT-AUTH-SECRET>"
  version     = "0.1.0"
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

output "my-terraform-template-id" {
  value = transloadit_template.my-terraform-template.id
}
```

Now on the cli, run:

```bash
terraform plan
```

## Contributing

### Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```bash
make build
# ...
$GOPATH/bin/terraform-provider-transloadit
# ...
```

In order to test the provider, you can simply run `make test`.

```bash
make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```bash
make testacc
```

### Releasing

Travis CI is set up to automatically build and release artifacts onto GitHub Releases upon pushing e.g. a "v0.2.0" tag to the master branch.
