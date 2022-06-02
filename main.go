package main

import (
	"fmt"
	"os"
	"strconv"
)

// Goal use main.go to build tf configurations and build terratest-plans for those configurations
// Configurations will be backed up and easily reverted to for testing purposes
// Build a UI that allows you to build tf configs and allow options to select pre-built tests to be applied for each config
// Once you have your sequence of configs and test cases for each config, you'll be able to run tests and view results

type Nodepool struct {
	Quantity int    `json:"quantity"`
	Etcd     string `json:"etcd"`
	Cp       string `json:"cp"`
	Wkr      string `json:"wkr"`
}

// Currently supports building tf module configurations

func main() {
	// Select from the following modules:
	// "aks", "rke1", "rke2"
	module := "rke2"

	// Customize your desired node pools
	// Update append() on line 75 to include desired node pools
	// Run `go run main.go` from root directory to build config
	pool1 := Nodepool{
		Quantity: 1,
		Etcd:     "true",
		Cp:       "true",
		Wkr:      "true",
	}

	pool2 := Nodepool{
		Quantity: 2,
		Etcd:     "false",
		Cp:       "true",
		Wkr:      "true",
	}

	pool3 := Nodepool{
		Quantity: 3,
		Etcd:     "true",
		Cp:       "false",
		Wkr:      "true",
	}

	pool4 := Nodepool{
		Quantity: 4,
		Etcd:     "true",
		Cp:       "true",
		Wkr:      "false",
	}

	pool5 := Nodepool{
		Quantity: 5,
		Etcd:     "false",
		Cp:       "false",
		Wkr:      "true",
	}

	pool6 := Nodepool{
		Quantity: 6,
		Etcd:     "false",
		Cp:       "true",
		Wkr:      "false",
	}

	pool7 := Nodepool{
		Quantity: 7,
		Etcd:     "true",
		Cp:       "true",
		Wkr:      "true",
	}

	var NodePools []Nodepool

	NodePools = append(NodePools, pool1, pool2, pool3, pool4, pool5, pool6, pool7)

	SetConfigTF(module, NodePools)

}

// end of func main()

// move function to package functions
func SetConfigTF(module string, nodePools []Nodepool) bool {
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
`

	f, err := os.Create("./modules/" + module + "/main.tf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", f)

	defer f.Close()

	switch {
	case module == "aks":
		config = config + `
resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
  azure_credential_config {
    client_id       = var.azure_client_id
	client_secret   = var.azure_client_secret
	subscription_id = var.azure_subscription_id
  }
}
`
		poolConfig := ``
		num := 1
		for _, pool := range nodePools {
			poolNum := strconv.Itoa(num)
			quantity := strconv.Itoa(pool.Quantity)
			poolConfig = poolConfig + `    node_pools {
      availability_zones = var.availability_zones
	  name = "pool` + poolNum + `"
	  count = ` + quantity + `
	  orchestrator_version = var.orchestrator_version
	  os_disk_size_gb = var.os_disk_size_gb
	  vm_size = var.vm_size
	}
`
			num = num + 1
		}
		config = config + `
resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  aks_config_v2 {
    cloud_credential_id = rancher2_cloud_credential.rancher2_cloud_credential.id
    resource_group = var.resource_group
    resource_location = var.resource_location
	dns_prefix = var.dns_prefix
	kubernetes_version = var.kubernetes_version
	network_plugin = var.network_plugin
` + poolConfig + `  }
}
`
		_, err = f.WriteString(config)

		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	case module == "rke1":
		config = config + `
resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
  amazonec2_credential_config {
    access_key = var.aws_access_key
    secret_key = var.aws_secret_key
    default_region = var.aws_region
  }
}
` + `
resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  rke_config {
	  kubernetes_version = var.kubernetes_version
	  network {
	    plugin = var.network_plugin
	  }
  }
}
	  ` + `
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
}`
		poolConfig := ``
		num := 1
		for _, pool := range nodePools {
			poolNum := strconv.Itoa(num)
			quantity := strconv.Itoa(pool.Quantity)
			poolConfig = poolConfig + `

resource "rancher2_node_pool" "pool` + poolNum + `" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool` + poolNum + `"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = ` + quantity + `
  control_plane    = ` + pool.Cp + `
  etcd             = ` + pool.Etcd + ` 
  worker           = ` + pool.Wkr + ` 
}`
			num = num + 1
		}
		config = config + poolConfig

		_, err = f.WriteString(config)

		if err != nil {
			fmt.Println(err)
			return false
		}
		return true

	case module == "rke2" || module == "k3s":
		config = config + `
data "rancher2_cloud_credential" "rancher2_cloud_credential" {
	name = var.cloud_credential_name
}
` + `
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
`
		poolConfig := ``
		num := 1
		for _, pool := range nodePools {
			poolNum := strconv.Itoa(num)
			quantity := strconv.Itoa(pool.Quantity)
			poolConfig = poolConfig + `    machine_pools {
	  name                         = "pool` + poolNum + `"
	  cloud_credential_secret_name = data.rancher2_cloud_credential.rancher2_cloud_credential.id
	  control_plane_role           = ` + pool.Cp + `
	  etcd_role                    = ` + pool.Etcd + `
	  worker_role                  = ` + pool.Wkr + `
	  quantity                     = ` + quantity + `
	  machine_config {
		kind = rancher2_machine_config_v2.rancher2_machine_config_v2.kind
		name = rancher2_machine_config_v2.rancher2_machine_config_v2.name
	  }
    }
`
			num = num + 1
		}
		config = config + `
resource "rancher2_cluster_v2" "rancher2_cluster_v2" {
  name                                     = var.cluster_name
  kubernetes_version                       = var.kubernetes_version
  enable_network_policy                    = var.enable_network_policy
  default_cluster_role_for_project_members = var.default_cluster_role_for_project_members
  rke_config {
` + poolConfig + `  }
}`

		_, err = f.WriteString(config)

		if err != nil {
			fmt.Println(err)
			return false
		}
		return true

	case module != "aks" || module != "rke1" || module != "rke2" || module != "k3s":
		fmt.Printf("\nModule does not exist; check for possible typo")
		return false
	default:
		return false
	}
}
