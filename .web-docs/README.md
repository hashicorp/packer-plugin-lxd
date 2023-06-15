The LXD plugin allows building containers for LXD, by starting an LXD container,
running provisioners within this container, then saveing the container
as an LXD image.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    lxd = {
      source  = "github.com/hashicorp/lxd"
      version = "~> 1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/hashicorp/lxd
```

### Components

#### Builders

- [LXD](/packer/integrations/hashicorp/lxd/latest/components/builder/lxd) - The LXD builder builds containers with LXD
  by starting a container, provisioning it, and exporting it as a tar.gz archive of the root file system.
