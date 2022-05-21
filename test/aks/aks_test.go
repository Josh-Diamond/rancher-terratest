package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/josh-diamond/rancher-terratest/functions"
)

func TestAKSDownStreamCluster(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../../modules/aks",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	url := terraform.Output(t, terraformOptions, "host_url")
	token := `Bearer ` + terraform.Output(t, terraformOptions, "bearer_token")
	name := terraform.Output(t, terraformOptions, "cluster_name_aks")
	id := functions.GetClusterID(url, name, token)

	
	expectedClusterName := name
	actualClusterName := functions.GetClusterName(url, id, token)
	assert.Equal(t, expectedClusterName, actualClusterName)

	expectedClusterNodeCount := 1
	actualClusterNodeCount := functions.GetClusterNodeCount(url, id, token)
	assert.Equal(t, expectedClusterNodeCount, actualClusterNodeCount)

	expectedClusterProvider := "aks"
	actualClusterProvider := functions.GetClusterProvider(url, id, token)
	assert.Equal(t, expectedClusterProvider, actualClusterProvider)

	expectedClusterState := "active"
	actualClusterState := functions.GetClusterState(url, id, token)
	assert.Equal(t, expectedClusterState, actualClusterState)
}
