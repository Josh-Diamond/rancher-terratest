package tests

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/josh-diamond/rancher-terratest/functions"
	"github.com/stretchr/testify/assert"
)

func TestAKSDownStreamCluster(t *testing.T) {
	t.Parallel()

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

	// Adds 3 node pools; each with 1 node
	updatedNodePools := functions.UpdateNodePoolsTF(actualClusterProvider, 3, 1, "")
	assert.Equal(t, updatedNodePools, true)

	terraformApplyUpdate()
	functions.WaitForActiveCLuster(url, name, token)
	time.Sleep(30 * time.Second)
	functions.WaitForActiveCLuster(url, name, token)

	expectedPostUpdate1TotalNodeCount := 4
	actualPostUpdate1NodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedPostUpdate1TotalNodeCount, actualPostUpdate1NodeCount)

	// Delete 3 node pools added above
	updatedNodePools2 := functions.UpdateNodePoolsTF(actualClusterProvider, 0, 0, "")
	assert.Equal(t, updatedNodePools2, true)

	terraformApplyUpdate()
	functions.WaitForActiveCLuster(url, name, token)
	time.Sleep(30 * time.Second)
	functions.WaitForActiveCLuster(url, name, token)

	actualPostUpdate2TotalNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedClusterNodeCount, actualPostUpdate2TotalNodeCount)

}
