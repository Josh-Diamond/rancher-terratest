package tests

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/josh-diamond/rancher-terratest/config"
	"github.com/josh-diamond/rancher-terratest/functions"
	"github.com/stretchr/testify/assert"
)

func TestAKSDownStreamCluster2(t *testing.T) {
	t.Parallel()

	config.BuildConfig1()
	result := functions.SetConfigTF(config.Aks, config.Config1)
	assert.Equal(t, true, result)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../../modules/aks",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	terraformApplyUpdate := func() {
		terraform.Apply(t, terraformOptions)
		terraform.Apply(t, terraformOptions)
	}

	url := terraform.Output(t, terraformOptions, "host_url")
	token := terraform.Output(t, terraformOptions, "token_prefix") + terraform.Output(t, terraformOptions, "token")
	name := terraform.Output(t, terraformOptions, "cluster_name")
	id := functions.GetClusterID(url, name, token)

	expectedClusterName := name
	actualClusterName := functions.GetClusterName(url, id, token)
	assert.Equal(t, expectedClusterName, actualClusterName)

	// TF output returns the value as type string, which will fail tests, as that's not the expected type from rancher server
	expectedClusterNodeCount := functions.OutputToInt(terraform.Output(t, terraformOptions, "config1_expected_node_count"))
	actualClusterNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedClusterNodeCount, actualClusterNodeCount)

	expectedClusterProvider := terraform.Output(t, terraformOptions, "config1_expected_provider")
	actualClusterProvider := functions.GetClusterProvider(url, id, token)
	assert.Equal(t, expectedClusterProvider, actualClusterProvider)

	expectedClusterState := terraform.Output(t, terraformOptions, "config1_expected_state")
	actualClusterState := functions.GetClusterState(url, id, token)
	assert.Equal(t, expectedClusterState, actualClusterState)

	expectedKubernetesVersion := terraform.Output(t, terraformOptions, "config1_expected_kubernetes_version")
	actualKubernetesVersion := functions.GetKubernetesVersion(url, id, token)
	assert.Equal(t, expectedKubernetesVersion, actualKubernetesVersion)

	expectedRancherServerVersion := terraform.Output(t, terraformOptions, "config1_expected_rancher_server_version")
	actualRancherServerVersion := functions.GetRancherServerVersion(url, token)
	assert.Equal(t, expectedRancherServerVersion, actualRancherServerVersion)

	// Builds + Sets Config2 + tests if successful
	config.BuildConfig2()
	result2 := functions.SetConfigTF(config.Aks, config.Config2)
	assert.Equal(t, result2, true)
	// TF Applies Config2
	terraformApplyUpdate()
	functions.WaitForActiveCLuster(url, name, token)
	time.Sleep(30 * time.Second)
	functions.WaitForActiveCLuster(url, name, token)
	// Test against Config2
	expectedConfig2NodeCount := terraform.Output(t, terraformOptions, "config2_expected_node_count")
	actualConfig2NodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedConfig2NodeCount, actualConfig2NodeCount)

	// Config3
	config.BuildConfig3()
	result3 := functions.SetConfigTF(config.Aks, config.Config3)
	assert.Equal(t, true, result3)

	terraformApplyUpdate()
	functions.WaitForActiveCLuster(url, name, token)
	time.Sleep(30 * time.Second)
	functions.WaitForActiveCLuster(url, name, token)

	expectedConfig3NodeCount := terraform.Output(t, terraformOptions, "config3_expected_node_count")
	actualConfig3TotalNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedConfig3NodeCount, actualConfig3TotalNodeCount)

}
