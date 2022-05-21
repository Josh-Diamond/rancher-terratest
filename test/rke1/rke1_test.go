package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/josh-diamond/rancher-terratest/functions"
	"github.com/stretchr/testify/assert"
)

func TestRke1DownStreamCluster(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../../modules/rke1",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	url := terraform.Output(t, terraformOptions, "host_url")
	token := terraform.Output(t, terraformOptions, "token_type") + terraform.Output(t, terraformOptions, "token")
	name := terraform.Output(t, terraformOptions, "cluster_name")
	functions.WaitForActiveCLuster(url, name, token)
	id := functions.GetClusterID(url, name, token)


	expectedClusterName := name
	actualClusterName := functions.GetClusterName(url, id, token)
	assert.Equal(t, expectedClusterName, actualClusterName)

	expectedClusterNodeCount := terraform.Output(t, terraformOptions, "expected_node_count")
	actualClusterNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedClusterNodeCount, actualClusterNodeCount)

	expectedClusterProvider := terraform.Output(t, terraformOptions, "expected_provider")
	actualClusterProvider := functions.GetClusterProvider(url, id, token)
	assert.Equal(t, expectedClusterProvider, actualClusterProvider)

	expectedClusterState := terraform.Output(t, terraformOptions, "expected_state")
	actualClusterState := functions.GetClusterState(url, id, token)
	assert.Equal(t, expectedClusterState, actualClusterState)

}

// RKE1 does not wait for cluster to provision before destroying.
// With RKE1, once the POST req is successful, terraform completes the job,
// runs tests pre-maturely while cluster is provisioning, then destroys cluster,
// failing all tests
//
// Solution: use WaitForActiveCluster() after provisioning and before test cases
//
// Additional thought: WaitForActiveCluster() might be useful when adding/deleting node pools or updating the cluster;

