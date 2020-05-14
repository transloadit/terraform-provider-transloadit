# Transloadit Terraform Provider

- Transloadit Docs: <https://transloadit.com/docs/>
- Transloadit Twitter: <https://twitter.com/transloadit>
- Transloadit Community: <https://community.transloadit.com/>
- Terraform Website: https://www.terraform.io
- Terraform Gitter: [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Terraform Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

<img src="https://transloadit.edgly.net/assets/images/artwork/logos-transloadit-default.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x

## Install the Provider

The recommended way to install the Transloadit Terraform Provider is to use binary distributions from the [Releases](https://github.com/transloadit/terraform-provider-transloadit/releases) page. The binaries are available for Linux, macOS (darwin) and Windows (all for the amd64 architecture).

Here's how to download and uncompress the latest release for your OS:

### Install the Provider on Linux

```bash
mkdir -p ~/.terraform.d/plugins/ && cd !$
curl -sSL https://github.com/transloadit/terraform-provider-transloadit/releases/download/v0.1.0/terraform-provider-transloadit_linux_amd64.tar.gz |tar xvz
cd -
```

### Install the Provider on macOS

```bash
mkdir -p ~/.terraform.d/plugins/ && cd !$
curl -sSL https://github.com/transloadit/terraform-provider-transloadit/releases/download/v0.1.0/terraform-provider-transloadit_darwin_amd64.tar.gz |tar xvz
cd -
```

### Install the Provider on Windows

[Download](https://github.com/transloadit/terraform-provider-transloadit/releases) and unpack the provider into `%APPDATA%\terraform.d\plugins`.

## Usage

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

Now on the CLI, run:

```bash
terraform init # only required on the first run after using a new Provider (version)
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

Travis CI is set up to automatically build and release artifacts onto GitHub Releases upon pushing e.g. a `v0.2.0` tag to the master branch.
