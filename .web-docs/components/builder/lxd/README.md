Type: `lxd`
Artifact BuilderId: `lxd`

The `lxd` Packer builder builds containers for LXD. The builder starts an LXD
container, runs provisioners within this container, then saves the container as
an LXD image.

The LXD builder requires a modern linux kernel and the `lxd` package. This
builder does not work with LXC.

## Basic Example

Below is a fully functioning example.

**HCL**

```hcl
source "lxd" "lxd-xenial" {
  image = "ubuntu-daily:xenial"
  output_image = "ubuntu-xenial"
  publish_properties {
    description = "Trivial repackage with Packer"
  }
}

build {
  sources = ["lxd.lxd-xenial"]
}
```

**JSON**

```json
{
  "builders": [
    {
      "type": "lxd",
      "name": "lxd-xenial",
      "image": "ubuntu-daily:xenial",
      "output_image": "ubuntu-xenial",
      "publish_properties": {
        "description": "Trivial repackage with Packer"
      }
    }
  ]
}
```

## Configuration Reference

### Required:

<!-- Code generated from the comments of the Config struct in builder/lxd/config.go; DO NOT EDIT MANUALLY -->

- `image` (string) - The source image to use when creating the build
  container. This can be a (local or remote) image (name or fingerprint).
  E.G. my-base-image, ubuntu-daily:x, 08fababf6f27, ...

<!-- End of code generated from the comments of the Config struct in builder/lxd/config.go; -->


  ~> Note: The builder may appear to pause if required to download a
  remote image, as they are usually 100-200MB. `/var/log/lxd/lxd.log` will
  mention starting such downloads.

### Optional:

<!-- Code generated from the comments of the Config struct in builder/lxd/config.go; DO NOT EDIT MANUALLY -->

- `output_image` (string) - The name of the output artifact. Defaults to
  name.

- `container_name` (string) - Container Name

- `publish_remote_name` (string) - The (optional) name of the LXD remote on which to publish the
  container image.

- `command_wrapper` (string) - Lets you prefix all builder commands, such as
  with ssh for a remote build host. Defaults to `{{.Command}}`; i.e. no
  wrapper.

- `profile` (string) - Profile

- `init_sleep` (string) - The number of seconds to sleep between launching
  the LXD instance and provisioning it; defaults to 3 seconds.

- `publish_properties` (map[string]string) - Pass key values to the publish
  step to be set as properties on the output image. This is most helpful to
  set the description, but can be used to set anything needed. See
  https://stgraber.org/2016/03/30/lxd-2-0-image-management-512/
  for more properties.

- `launch_config` (map[string]string) - List of key/value pairs you wish to
  pass to lxc launch via --config. Defaults to empty.

- `virtual_machine` (bool) - Create LXD virtual-machine image on hosts running LXD 4.0 and above; defaults to false for container image

- `skip_publish` (bool) - Skip execute `lxc publish`; defaults to false

<!-- End of code generated from the comments of the Config struct in builder/lxd/config.go; -->
