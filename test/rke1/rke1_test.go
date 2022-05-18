package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRke1DownSteamCluster(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../../modules/rke1",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	

	expectedClusterName := "tf-rke1-test"
	actualClusterName := terraform.Output(t, terraformOptions, "cluster_name_rke1")
	assert.Equal(t, expectedClusterName, actualClusterName)

}

// RKE1 does not wait for cluster to provision before destroying.  
// With RKE1, once the POST req is successful, terraform completes the job, 
// runs tests pre-maturely while cluster is provisioning, and destroys cluster, 
// failing all tests
// Thought: create a conditional loop that listens on cluster object endpoint immediately after cluster creation, 
// and only continues when clusters state is active