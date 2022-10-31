locals {
  instances = {
    "micro-node" = {
      cpu = 2
      ram = 2
    }
    "macro-node" = {
      cpu = 4
      ram = 4
    }
  }

  #Вводим переменную, которая определяет, находимся ли мы в workspace prod
  is_prod = terraform.workspace == "prod" ? true : false
}

resource "yandex_compute_instance" "netology_node_07_terraform_for_each" {
  #Если мы не находимся в workspace prod, то передаем пустой список
  for_each = local.is_prod ? local.instances : { }

  name                      = "node-07-terraform-${each.key}"
  zone                      = "ru-central1-b"
  hostname                  = "node-07-terraform.netology.cloud"
  allow_stopping_for_update = true

  resources {
    cores  = each.value["cpu"]
    memory = each.value["ram"]
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
}
