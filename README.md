Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/Etienne-Carriere/terraform-provider-transloadit`

```sh
$ mkdir -p $GOPATH/src/github.com/Etienne-Carriere; cd $GOPATH/src/github.com/Etienne-Carriere
$ git clone https://github.com/Etienne-Carriere/terraform-provider-transloadit.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-transloadit
$ make build
```

Installing the provider
-----------------------

### Installation from binaries (recommended)

The recommended way to install *terraform-provider-transloadit* is use the binary
distributions from the [Releases](https://github.com/transloadit/terraform-provider-transloadit/releases) page. The packages are available for Linux , macOS (darwin) and Windows (on amd64 architecture)

Download and uncompress the latest release for your OS. This example uses the linux binary.

```sh
$ wget https://github.com/transloadit/terraform-provider-transloadit/releases/download/v0.1.0/terraform-provider-transloadit_linux_amd64.tar.gz
$ tar -xvf terraform-provider-transloadit*.tar.gz
```

Now copy the binary to the Terraform's plugins folder (if this is your first plugin maybe it isn't present):

```sh
$ mkdir -p ~/.terraform.d/plugins/
$ mv terraform-provider-transloadit* ~/.terraform.d/plugins/
```

On windows, you have to put the provider in `%APPDATA%\terraform.d\plugins` 

Using the provider
----------------------
See in the website directory

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-transloadit
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
