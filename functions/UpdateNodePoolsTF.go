package functions

import (
	"fmt"
	"strconv"
)

// This page is has string literals formatted specifically which should not be modified; W.I.P.
func UpdateNodePoolsTF(provider string, poolcount int, nodecount int, role string ) bool {
	switch {
	case provider == "aks":
		if poolcount == 0 {
			return false
		}
		poolConfig := ``
		name := 2
		for i := 1; i <= poolcount; i++ {
		poolNum := strconv.Itoa(name)
		count := strconv.Itoa(nodecount)
		poolConfig = poolConfig + `node_pools {
		availability_zones = var.availability_zones
		name = "pool`+ poolNum + `"
		count = `+ count +`
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
	}
	` + poolConfig + `
  }
}
	`
		  UpdateConfigTF(config, provider)
		  return true
	case provider == "rke":
		fmt.Println(provider)
	case provider == "rke2" || provider == "k3s":
		fmt.Println(provider)
	default:
		return false

	}

	return false
}