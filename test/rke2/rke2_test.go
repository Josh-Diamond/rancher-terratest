package test

import (
	"testing"
	"github.com/josh-diamond/rancher-terratest/functions"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRke2DownSteamCluster(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../../modules/rke2",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	
	id := functions.GetClusterID("URL_here", "Bearer_token_here")

	expectedClusterNameTF := "expected-name"
	actualClusterNameTF := terraform.Output(t, terraformOptions, "cluster_name_rke2")
	assert.Equal(t, expectedClusterNameTF, actualClusterNameTF)

	expectedClusterName := "expected-name"
	actualClusterName := functions.GetClusterName("URL_HERE", id, "Bearer_token_here")
	assert.Equal(t, expectedClusterName, actualClusterName)

	expectedClusterNodeCount := 1
	actualClusterNodeCount := functions.GetClusterNodeCount("URL_HERE", id, "Bearer_token_here")
	assert.Equal(t, expectedClusterNodeCount, actualClusterNodeCount)

	expectedClusterProvider := "expected-provider"
	actualClusterProvider := functions.GetClusterProvider("URL_HERE", id, "Bearer_token_here")
	assert.Equal(t, expectedClusterProvider, actualClusterProvider)

	expectedClusterState := "expected-state"
	actualClusterState := functions.GetClusterState("URL_HERE", id, "Bearer_token_here")
	assert.Equal(t, expectedClusterState, actualClusterState)

}
