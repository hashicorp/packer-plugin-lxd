The incus plugin allows building containers for incus, by starting an incus container,
running provisioners within this container, then saveing the container
as an incus image.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    incus = {
      source  = "github.com/bketelsen/incus"
      version = "~> 1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/bketelsen/incus
```

### Components

#### Builders

- [incus](/packer/integrations/bketelsen/incus/latest/components/builder/incus) - The incus builder builds containers with incus
  by starting a container, provisioning it, and exporting it as a tar.gz archive of the root file system.
