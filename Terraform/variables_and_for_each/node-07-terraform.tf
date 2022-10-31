locals {
  compute_instance_cores = {
    stage = 2
    prod = 4
  }

  compute_instance_ram = {
    stage = 2
    prod = 4
  }

  compute_instance_count = {
    stage = 1
    prod = 2
  }
}

resource "yandex_compute_instance" "netology_node_07_terraform" {
  name                      = "node-07-terraform"
  zone                      = "ru-central1-b"
  hostname                  = "node-07-terraform.netology.cloud"
  allow_stopping_for_update = true
  count = local.compute_instance_count[terraform.workspace]

  resources {
    cores  = local.compute_instance_cores[terraform.workspace]
    memory = local.compute_instance_ram[terraform.workspace]
  }

  boot_disk {
    initialize_params {
      image_id    = "${var.centos-7-base}"
      name        = "root-node01"
      type        = "network-nvme"
      size        = "20"
    }
  }

  network_interface {
    subnet_id = "${yandex_vpc_subnet.netology_07_terraform.id}"
    nat       = true
  }

  metadata = {
    ssh-keys = "centos:${file("~/.ssh/id_rsa.pub")}"
  }

  lifecycle {
    create_before_destroy = true
  }
}
