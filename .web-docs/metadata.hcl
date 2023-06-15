# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "LXD"
  description = "The LXD plugin can be used with HashiCorp Packer to create OCI images with LXD."
  identifier = "packer/hashicorp/lxd"
  component {
    type = "builder"
    name = "LXD"
    slug = "lxd"
  }
}
