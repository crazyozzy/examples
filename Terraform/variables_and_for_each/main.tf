terraform {
  required_providers {
    yandex = {
      source = "yandex-cloud/yandex"
    }
  }

  backend "s3" {
  endpoint   = "storage.yandexcloud.net"
  bucket     = "netology-terraform"
  region     = "ru-central1"
  key        = "devops-netology.tfstate"
#  access_key = "<идентификатор статического ключа>"
#  secret_key = "<секретный ключ>"

  skip_region_validation      = true
  skip_credentials_validation = true
  }

  required_version = ">= 0.13"
}

provider "yandex" {
  zone = "ru-central1-b"
}
