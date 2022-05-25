package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/josh-diamond/rancher-terratest/functions"
	"github.com/stretchr/testify/assert"
)

func TestRke2DownStreamCluster(t *testing.T) {

const config = `terraform {
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
		machine_pools {
			name                         = "pool2"
			cloud_credential_secret_name = data.rancher2_cloud_credential.rancher2_cloud_credential.id
			control_plane_role           = false
			etcd_role                    = false
			worker_role                  = true
			quantity                     = 3
			machine_config {
			kind = rancher2_machine_config_v2.rancher2_machine_config_v2.kind
			name = rancher2_machine_config_v2.rancher2_machine_config_v2.name
			}
		}
	}
}

`	

	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../../modules/rke2",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	url := terraform.Output(t, terraformOptions, "host_url")
	token := terraform.Output(t, terraformOptions, "token_prefix") + terraform.Output(t, terraformOptions, "token")
	name := terraform.Output(t, terraformOptions, "cluster_name")
	id := functions.GetClusterID(url, name, token)

	expectedClusterName := name
	actualClusterName := functions.GetClusterName(url, id, token)
	assert.Equal(t, expectedClusterName, actualClusterName)

	expectedClusterNodeCount := functions.OutputToInt(terraform.Output(t, terraformOptions, "expected_node_count"))
	actualClusterNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedClusterNodeCount, actualClusterNodeCount)

	expectedClusterProvider := terraform.Output(t, terraformOptions, "expected_provider")
	actualClusterProvider := functions.GetClusterProvider(url, id, token)
	assert.Equal(t, expectedClusterProvider, actualClusterProvider)

	expectedClusterState := terraform.Output(t, terraformOptions, "expected_state")
	actualClusterState := functions.GetClusterState(url, id, token)
	assert.Equal(t, expectedClusterState, actualClusterState)

	expectedKubernetesVersion := terraform.Output(t, terraformOptions, "expected_kubernetes_version")
	actualKubernetesVersion := functions.GetKubernetesVersion(url, id, token)
	assert.Equal(t, expectedKubernetesVersion, actualKubernetesVersion)

	expectedRancherServerVersion := terraform.Output(t, terraformOptions, "expected_rancher_server_version")
	actualRancherServerVersion := functions.GetRancherServerVersion(url, token)
	assert.Equal(t, expectedRancherServerVersion, actualRancherServerVersion)

	functions.UpdateConfigTF(config, "rke2")
	terraform.Apply(t, terraformOptions)

	functions.WaitForActiveCLuster(url, name, token)

	expectedClusterNodeCountPost := 4
	actualClusterNodeCountPost := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedClusterNodeCountPost, actualClusterNodeCountPost)

}
