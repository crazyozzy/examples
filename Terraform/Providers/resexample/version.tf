terraform {
  required_providers {
    resexample = {
      version = "~> 1.0.0"
      source  = "k11s.ru/custom-provider/resexample"
    }
  }
}

provider "resexample" {
  api_version = "resource"
  hostname = "localhost:8080"
}