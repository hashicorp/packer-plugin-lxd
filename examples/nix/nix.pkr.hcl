source "incus" "nix" {
  image = "images:ubuntu/jammy"
  output_image = "ubuntu-nix"
  reuse = true
}

build {
  sources = ["incus.nix"]
  provisioner "file" {
    source = "examples/common/90-incus"
    destination = "/tmp/90-incus"
  }
  provisioner "shell" {
    scripts = [
      "examples/common/packages.sh",
      "examples/common/user.sh",
      "examples/common/sudoers.sh",
      "examples/common/nix.sh",
    ]
  }

}

