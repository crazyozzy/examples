resource "yandex_vpc_network" "netology_07_terraform" {
  name = "net"
}

resource "yandex_vpc_subnet" "netology_07_terraform" {
  name = "subnet"
  zone           = "ru-central1-b"
  network_id     = "${yandex_vpc_network.netology_07_terraform.id}"
  v4_cidr_blocks = ["192.168.101.0/24"]
}
