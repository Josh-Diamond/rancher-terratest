terraform {
  required_providers {
    rancher2 = {
      source  = "rancher/rancher2"
      version = "1.22.2"
    }
  }
}

provider "rancher2" {
  api_url   = var.rancher_api_url
  token_key = var.rancher_admin_bearer_token
}

# resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
#   name = var.cloud_credential_name
#   amazonec2_credential_config {
#     access_key = var.aws_access_key
#     secret_key = var.aws_secret_key
#       default_region = var.aws_region
#   }
# }

data "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
}

resource "rancher2_machine_config_v2" "rancher2_machine_config_v2" {
  generate_name = var.machine_config_name
  amazonec2_config {
    ami            = var.aws_ami
    region         = var.aws_region
    security_group = [var.aws_security_group_name]
    subnet_id      = var.aws_subnet_id
    vpc_id         = var.aws_vpc_id
    zone           = var.aws_zone_letter
  }
}

# RKE2/k3s is determined by the k8s version that is used below
resource "rancher2_cluster_v2" "rancher2_cluster_v2" {
  name                                     = var.cluster_name
  kubernetes_version                       = var.kubernetes_version
  enable_network_policy                    = var.enable_network_policy
  default_cluster_role_for_project_members = var.default_cluster_role_for_project_members
  rke_config {
    machine_pools {
      name                         = "pool1"
      cloud_credential_secret_name = data.rancher2_cloud_credential.rancher2_cloud_credential.id
      control_plane_role           = true
      etcd_role                    = true
      worker_role                  = true
      quantity                     = 1
      machine_config {
        kind = rancher2_machine_config_v2.rancher2_machine_config_v2.kind
        name = rancher2_machine_config_v2.rancher2_machine_config_v2.name
      }
    }
  }
}

# To-do: Create resource for cloud credentials instead of pulling data from rancher server.  Neccessary to fully automate
# │ ERROR
# │ When expanding the plan for rancher2_cluster_v2.rancher2_cluster_v2 to include new values learned so far during
# │ apply, provider "registry.terraform.io/rancher/rancher2" produced an invalid new value for
# │ .rke_config[0].machine_pools[0].cloud_credential_secret_name: was cty.StringVal(""), but now
# │ cty.StringVal("cattle-global-data:REDACTED").
# │
# │ This is a bug in the provider, which should be reported in the provider's own issue tracker.