# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "incus"
  description = "The incus plugin can be used with HashiCorp Packer to create OCI images with incus."
  identifier = "packer/bketelsen/incus"
  component {
    type = "builder"
    name = "incus"
    slug = "incus"
  }
}
