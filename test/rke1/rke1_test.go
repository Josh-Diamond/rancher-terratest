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
