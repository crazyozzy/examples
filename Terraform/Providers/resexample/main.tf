data "resexample_create" "vm_1_name" {
  region = "MSK"
  resource_type = "server"
}

data "resexample_details" "vm_1_details" {
  name = data.resexample_create.vm_1_name.name
}

output "vm_1_name" {
  value = data.resexample_create.vm_1_name.name
}

data "resexample_create" "vm_2_name" {
  region = "Brazil"
  resource_type = "cocks_and_whisky"
}

data "resexample_details" "vm_2_details" {
  name = data.resexample_create.vm_2_name.name
}

output "vm_2_name" {
  value = data.resexample_create.vm_2_name.name
}