package functions

import "strconv"

// This page is has string literals specifically formatted which should not be modified;
func UpdateNodePoolsTF(provider string, quantity int, nodecount int, role string) bool {
	switch {
	case provider == "aks":
		if quantity == 0 {
			config := `terraform {
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
			  
resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
  azure_credential_config {
    client_id       = var.azure_client_id
	client_secret   = var.azure_client_secret
	subscription_id = var.azure_subscription_id
  }
}
			  
resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  aks_config_v2 {
    cloud_credential_id = rancher2_cloud_credential.rancher2_cloud_credential.id
	resource_group = var.resource_group
	resource_location = var.resource_location
	dns_prefix = var.dns_prefix
	kubernetes_version = var.kubernetes_version
	network_plugin = var.network_plugin
	node_pools {
	  availability_zones = var.availability_zones
	  name = "pool1"
	  count = 1
	  orchestrator_version = var.orchestrator_version
	  os_disk_size_gb = var.os_disk_size_gb
	  vm_size = var.vm_size
	}
  }
}
`
			UpdateConfigTF(config, provider)
			return true
		}
		if quantity != 0 {

			poolConfig := ``
			name := 2
			for i := 1; i <= quantity; i++ {
				poolNum := strconv.Itoa(name)
				count := strconv.Itoa(nodecount)
				poolConfig = poolConfig + `
	node_pools {
		availability_zones = var.availability_zones
		name = "pool` + poolNum + `"
		count = ` + count + `
		orchestrator_version = var.orchestrator_version
		os_disk_size_gb = var.os_disk_size_gb
		vm_size = var.vm_size
	}`
				name = name + 1
			}
			config := `terraform {
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

resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
  azure_credential_config {
	client_id       = var.azure_client_id
	client_secret   = var.azure_client_secret
	subscription_id = var.azure_subscription_id
  }
}

resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  aks_config_v2 {
	cloud_credential_id = rancher2_cloud_credential.rancher2_cloud_credential.id
	resource_group = var.resource_group
	resource_location = var.resource_location
	dns_prefix = var.dns_prefix
	kubernetes_version = var.kubernetes_version
	network_plugin = var.network_plugin
	node_pools {
		availability_zones = var.availability_zones
		name = "pool1"
		count = 1
		orchestrator_version = var.orchestrator_version
		os_disk_size_gb = var.os_disk_size_gb
		vm_size = var.vm_size
	}` + poolConfig + `
  }
}
	`
			UpdateConfigTF(config, provider)
			return true
		}
	case provider == "rke":
		if quantity == 0 {
			config := `terraform {
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
						
resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
  amazonec2_credential_config {
    access_key = var.aws_access_key
    secret_key = var.aws_secret_key
    default_region = var.aws_region
  }
}
						
# Create a new rancher2 RKE Cluster
resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  rke_config {
    kubernetes_version = var.kubernetes_version
	network {
	  plugin = var.network_plugin
	}
  }
}
			  
# Create a new rancher2 Node Template
resource "rancher2_node_template" "rancher2_node_template" {
  name = var.node_template_name
  amazonec2_config {
    access_key     = var.aws_access_key
	secret_key     = var.aws_secret_key
	ami            = var.aws_ami_w_docker
	region         = var.aws_region
	security_group = [var.aws_security_group_name]
	subnet_id      = var.aws_subnet_id
	vpc_id         = var.aws_vpc_id
	zone           = var.aws_zone_letter
	root_size      = var.aws_root_size
	instance_type  = var.aws_instance_type
  }
}
			  
# Create a new rancher2 Node Pool
resource "rancher2_node_pool" "pool1" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool1"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 1
  control_plane    = true 
  etcd             = true 
  worker           = true 
}
`
			UpdateConfigTF(config, provider)
			return true
		}

		if quantity != 0 {

			poolConfig := ``
			name := 2
			for i := 1; i <= quantity; i++ {
				poolNum := strconv.Itoa(name)
				count := strconv.Itoa(nodecount)
				etcd := "false"
				cp := "false"
				wkr := "false"
				if role == "etcd" {
					etcd = `true`
				}
				if role == "controlplane" {
					cp = `true`
				}
				if role == "wkr" {
					wkr = `true`
				}
				if role == "all" {
					etcd = `true`
					cp = `true`
					wkr = `true`
				}
				poolConfig = poolConfig + `
resource "rancher2_node_pool" "pool` + poolNum + `" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool` + poolNum + `"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = ` + count + `
  control_plane    = ` + cp + ` 
  etcd             = ` + etcd + ` 
  worker           = ` + wkr + ` 
}
		  `
				name = name + 1
			}
			config := `terraform {
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
		  
resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
  amazonec2_credential_config {
    access_key = var.aws_access_key
    secret_key = var.aws_secret_key
    default_region = var.aws_region
  }
}
		  
# Create a new rancher2 RKE Cluster
resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  rke_config {
    kubernetes_version = var.kubernetes_version
	network {
	  plugin = var.network_plugin
	}
  }
}

# Create a new rancher2 Node Template
resource "rancher2_node_template" "rancher2_node_template" {
  name = var.node_template_name
  amazonec2_config {
    access_key     = var.aws_access_key
	secret_key     = var.aws_secret_key
	ami            = var.aws_ami_w_docker
	region         = var.aws_region
	security_group = [var.aws_security_group_name]
	subnet_id      = var.aws_subnet_id
	vpc_id         = var.aws_vpc_id
	zone           = var.aws_zone_letter
	root_size      = var.aws_root_size
	instance_type  = var.aws_instance_type
  }
}

# Create a new rancher2 Node Pool
resource "rancher2_node_pool" "pool1" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool1"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 1
  control_plane    = true 
  etcd             = true 
  worker           = true 
}
		  ` + poolConfig + `
	`
			UpdateConfigTF(config, provider)
			return true
		}
	case provider == "rke2" || provider == "k3s":
		if quantity == 0 {
			config := `terraform {
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
						
resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
  amazonec2_credential_config {
    access_key = var.aws_access_key
    secret_key = var.aws_secret_key
    default_region = var.aws_region
  }
}
						
# Create a new rancher2 RKE Cluster
resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  rke_config {
    kubernetes_version = var.kubernetes_version
	network {
	  plugin = var.network_plugin
	}
  }
}
			  
# Create a new rancher2 Node Template
resource "rancher2_node_template" "rancher2_node_template" {
  name = var.node_template_name
  amazonec2_config {
    access_key     = var.aws_access_key
	secret_key     = var.aws_secret_key
	ami            = var.aws_ami_w_docker
	region         = var.aws_region
	security_group = [var.aws_security_group_name]
	subnet_id      = var.aws_subnet_id
	vpc_id         = var.aws_vpc_id
	zone           = var.aws_zone_letter
	root_size      = var.aws_root_size
	instance_type  = var.aws_instance_type
  }
}
			  
# Create a new rancher2 Node Pool
resource "rancher2_node_pool" "pool1" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool1"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 1
  control_plane    = true 
  etcd             = true 
  worker           = true 
}
`
			UpdateConfigTF(config, provider)
			return true
		}

		if quantity != 0 {

			poolConfig := ``
			name := 2
			for i := 1; i <= quantity; i++ {
				poolNum := strconv.Itoa(name)
				count := strconv.Itoa(nodecount)
				etcd := "false"
				cp := "false"
				wkr := "false"
				if role == "etcd" {
					etcd = `true`
				}
				if role == "controlplane" {
					cp = `true`
				}
				if role == "wkr" {
					wkr = `true`
				}
				if role == "all" {
					etcd = `true`
					cp = `true`
					wkr = `true`
				}
				poolConfig = poolConfig + `
	machine_pools {
	  name                         = "pool` + poolNum + `"
	  cloud_credential_secret_name = data.rancher2_cloud_credential.rancher2_cloud_credential.id
	  control_plane_role           = ` + cp + `
	  etcd_role                    = ` + etcd + `
	  worker_role                  = ` + wkr + `
	  quantity                     = ` + count + `
	  machine_config {
	    kind = rancher2_machine_config_v2.rancher2_machine_config_v2.kind
		name = rancher2_machine_config_v2.rancher2_machine_config_v2.name
	  }
	}`
				name = name + 1
			}
			config := `terraform {
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
		` + poolConfig + `
  }
}
`
			UpdateConfigTF(config, provider)
			return true
		}
	default:
		return false

	}

	return false
}
