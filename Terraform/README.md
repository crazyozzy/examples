## Local plugin initialization

При использовании [инструкции](https://developer.hashicorp.com/terraform/tutorials/providers/provider-use?in=terraform%2Fproviders) по установке учебного провайдера hashicups столкнулся с ошибкой:

```
Initializing the backend...

Initializing provider plugins...
- Finding hashicorp.com/edu/hashicups versions matching "~> 0.3.1"...
╷
│ Error: Invalid provider registry host
│
│ The host "hashicorp.com" given in in provider source address "hashicorp.com/edu/hashicups" does not offer a Terraform
│ provider registry.
╵
```

Для решения проблемы необходимо прописать в файл ".terraformrc" (может находится в ~ или в директории с проектом)
следующие строки:
```hcl
provider_installation {
  filesystem_mirror {
    path    = "/home/<your_user_name>/.terraform.d/plugins"
  }
  direct {
    exclude = ["hashicorp.com/edu/*"]
  }
}
```

## Add Yandex mirror

Для добавления зеркала Yandex, необходимо прописать в файл ".terraformrc" (может находится в ~ или в директории с проектом)
следующие строки:
```hcl
provider_installation {
  network_mirror {
    url = "https://terraform-mirror.yandexcloud.net/"
    include = ["registry.terraform.io/*/*"]
  }
  direct {
    exclude = ["registry.terraform.io/*/*"]
  }
}
```