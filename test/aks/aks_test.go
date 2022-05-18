package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAKSDownSteamCluster(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../../modules/aks",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)


	expectedClusterName := "tf-aks"
	actualClusterName := terraform.Output(t, terraformOptions, "cluster_name_aks")
	assert.Equal(t, expectedClusterName, actualClusterName)

}
