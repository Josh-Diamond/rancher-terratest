package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/josh-diamond/rancher-terratest/config"
	"github.com/josh-diamond/rancher-terratest/functions"
	"github.com/stretchr/testify/assert"
)

func TestRke1DownStreamCluster(t *testing.T) {
	t.Parallel()

	config.BuildNodePools1()
	config1 := functions.SetConfigTF(config.Rke1, config.Config1RKE1K8sVersion, config.NodePools1)
	assert.Equal(t, true, config1)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../../modules/rke1",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	url := terraform.Output(t, terraformOptions, "host_url")
	token := terraform.Output(t, terraformOptions, "token_prefix") + terraform.Output(t, terraformOptions, "token")
	name := terraform.Output(t, terraformOptions, "cluster_name")
	functions.WaitForActiveCLuster(url, name, token)
	id := functions.GetClusterID(url, name, token)

	expectedClusterName := name
	actualClusterName := functions.GetClusterName(url, id, token)
	assert.Equal(t, expectedClusterName, actualClusterName)

	config1ExpectedNodeCount := functions.OutputToInt(terraform.Output(t, terraformOptions, "config1_expected_node_count"))
	config1ActualNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, config1ExpectedNodeCount, config1ActualNodeCount)

	config1ExpectedProvider := terraform.Output(t, terraformOptions, "config1_expected_provider")
	config1ActualProvider := functions.GetClusterProvider(url, id, token)
	assert.Equal(t, config1ExpectedProvider, config1ActualProvider)

	config1ExpectedState := terraform.Output(t, terraformOptions, "config1_expected_state")
	config1ActualState := functions.GetClusterState(url, id, token)
	assert.Equal(t, config1ExpectedState, config1ActualState)

	config1ExpectedKubernetesVersion := terraform.Output(t, terraformOptions, "config1_expected_kubernetes_version")
	config1ActualKubernetesVersion := functions.GetKubernetesVersion(url, id, token)
	assert.Equal(t, config1ExpectedKubernetesVersion, config1ActualKubernetesVersion)

	config1ExpectedRancherServerVersion := terraform.Output(t, terraformOptions, "config1_expected_rancher_server_version")
	config1ActualRancherServerVersion := functions.GetRancherServerVersion(url, token)
	assert.Equal(t, config1ExpectedRancherServerVersion, config1ActualRancherServerVersion)

	// Config2
	config.BuildNodePools2()
	config2 := functions.SetConfigTF(config.Rke1, config.Config2RKE1K8sVersion, config.NodePools2)
	assert.Equal(t, true, config2)

	terraform.Apply(t, terraformOptions)
	functions.WaitForActiveCLuster(url, name, token)

	config2ExpectedNodeCount := functions.OutputToInt(terraform.Output(t, terraformOptions, "config2_expected_node_count"))
	config2ActualNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, config2ExpectedNodeCount, config2ActualNodeCount)

	config2ExpectedKubernetesVersion := terraform.Output(t, terraformOptions, "config2_expected_kubernetes_version")
	config2ActualKubernetesVersion := functions.GetKubernetesVersion(url, id, token)
	assert.Equal(t, config2ExpectedKubernetesVersion, config2ActualKubernetesVersion)

	// config3
	config.BuildNodePools3()
	config3 := functions.SetConfigTF(config.Rke1, config.Config2RKE1K8sVersion, config.NodePools3)
	assert.Equal(t, true, config3)

	terraform.Apply(t, terraformOptions)
	functions.WaitForActiveCLuster(url, name, token)

	config3ExpectedNodeCount := functions.OutputToInt(terraform.Output(t, terraformOptions, "config3_expected_node_count"))
	config3ActualNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, config3ExpectedNodeCount, config3ActualNodeCount)

}
