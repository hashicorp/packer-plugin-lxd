source "incus" "jammy" {
  image = "images:ubuntu/jammy"
  output_image = "ubuntu-jammy"
  reuse = true
}

build {
  sources = ["incus.jammy"]
  provisioner "shell" {
    inline = [
      "apt-get update",
      "apt-get -y install openssh-server curl wget",
      ]
  }

}

