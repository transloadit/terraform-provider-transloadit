[![Build Status](https://travis-ci.org/transloadit/terraform-provider-transloadit.svg?branch=master)](https://travis-ci.org/transloadit/terraform-provider-transloadit)

# Transloadit Terraform Provider

<img align="right" src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" height="82px">

- Terraform Website: https://www.terraform.io
- Terraform Gitter: [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Terraform Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img align="right" src="https://transloadit.edgly.net/assets/images/artwork/logos-transloadit-default.svg" height="130px">

- Transloadit Docs: <https://transloadit.com/docs/>
- Transloadit Twitter: <https://twitter.com/transloadit>
- Transloadit Community: <https://community.transloadit.com/>

## Intro

Transloadit is a file processing service for companies. Developers describe desired outputs with a declarative JSON recipe to encode video, resize and recognize images, normalize audio, thumbnail documents, screenshot websites, and so much more.

Terraform is a tool to provision infrastructure as code. Developers describe desired infrastructure (such as webservers, loadbalancers, storage buckets) with a declarative HCL to launch it.

The Transloadit Terraform Provider lets you launch Transloadit recipes (Templates) with Terraform. This way you can orchestrate your media processing pipelines declaratively from a single place, and make sure it works well in conjunction with other infrastructure. For instance, you may want to use Transloadit as the media conversion engine for your not-so Smart CDN, or make sure required storage  buckets are set up in a compatible way.

## Install

You'll first need [Terraform](https://www.terraform.io/downloads.html) 0.12.x installed on your system as a prerequisite.

The recommended way to install the Transloadit Terraform Provider is to use binary distributions from the [Releases](https://github.com/transloadit/terraform-provider-transloadit/releases) page. The binaries are available for Linux, macOS (darwin) and Windows (all for the amd64 architecture).

Here's how to download and extract the latest release for your OS:

### Linux

```bash
mkdir -p ~/.terraform.d/plugins/ && cd !$
curl -sSL https://github.com/transloadit/terraform-provider-transloadit/releases/download/v0.1.0/terraform-provider-transloadit_linux_amd64.tar.gz |tar xvz
cd -
```

### macOS

```bash
mkdir -p ~/.terraform.d/plugins/ && cd !$
curl -sSL https://github.com/transloadit/terraform-provider-transloadit/releases/download/v0.1.0/terraform-provider-transloadit_darwin_amd64.tar.gz |tar xvz
cd -
```

### Windows

[Download](https://github.com/transloadit/terraform-provider-transloadit/releases) and unpack the provider into `%APPDATA%\terraform.d\plugins`.

## Usage

Here's a quick example. More detailed instructions can be found in the [website directory](./website/). If you intend to use multiple Apps withing Transloadit, check out the [examples](./examples/multi-app) how to structure that with Modules.

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

You'll need to copy `env.example.sh` to `env.sh` and add real Transloadit Credentials. You can then source that and run:

```bash
make testacc
```

*Note:* Acceptance tests create real resources, and often cost money to run.

### Releasing

Travis CI is set up to automatically build and release artifacts onto GitHub Releases upon pushing e.g. a `v0.2.0` tag to the master branch.
