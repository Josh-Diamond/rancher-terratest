package test

import (
	"testing"
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
	

	expectedClusterName := "tf-rke2-k3s"
	actualClusterName := terraform.Output(t, terraformOptions, "cluster_name_rke2")
	assert.Equal(t, expectedClusterName, actualClusterName)

}
