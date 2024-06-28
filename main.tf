terraform {
  required_providers {
    truenas = {
      source  = "registry.terraform.io/jdella/truenas"
    }
  }
}

variable TRUENAS_API_KEY {
  type        = string
}

variable TRUENAS_HOST_NAME {
  type        = string
}

provider "truenas" {
  api_key = var.TRUENAS_API_KEY
  base_url = "http://${var.TRUENAS_HOST_NAME}/api/v2.0"
}

# data "truenas_dataset" "dataset" {
#   dataset_id = "della-pool-rust/backup"
# } 

# output test {
#   value       = data.truenas_dataset.dataset
# }

data "truenas_services" "svc" {
}

data "truenas_service" "svc_breakdown" {
    for_each = {for element in data.truenas_services.svc.ids: element => element }
    service_id = each.value
}

output test2 {
  value       = data.truenas_service.svc_breakdown
}

# resource "truenas_user" "tu" {
#   email             = "jake.della5@gmail.com"
#   full_name         = "JDUrp"
#   create_group      = true
#   name              = "Jake_Della33"
#   home_mode         = "0775"
#   smb               = false
#   sudo              = false
#   password_disabled = false
#   password = "Foobar"
# }