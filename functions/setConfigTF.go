package functions

import (
	"fmt"
	"os"
	"strconv"

	"github.com/josh-diamond/rancher-terratest/components"
	"github.com/josh-diamond/rancher-terratest/models"
)

func SetConfigTF(module string, nodePools []models.Nodepool) bool {

	config := components.RequiredProviders + components.Provider

	f, err := os.Create("./modules/" + module + "/main.tf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", f)

	defer f.Close()

	switch {
	case module == "aks":
		config = config + components.AzureCloudCredentials
		poolConfig := ``
		num := 1
		for _, pool := range nodePools {
			poolNum := strconv.Itoa(num)
			quantity := strconv.Itoa(pool.Quantity)
			if pool.Etcd != "true" && pool.Cp != "true" && pool.Wkr != "true" {
				fmt.Println("")
				fmt.Printf(`No roles selected for pool` + poolNum + `; at least one role is required`)
			}
			if pool.Quantity <= 0 {
				fmt.Println("")
				fmt.Println(`Invalid quantity specified for pool` + poolNum + `. Quantity must be greater than 0`)
			}
			poolConfig = poolConfig + components.AKSNodePoolPrefix + poolNum + components.AKSNodePoolBody + quantity + components.AKSNodePoolSuffix
			num = num + 1
		}
		config = config + components.AKSClusterPrefix + poolConfig + components.AKSClusterSuffix
		_, err = f.WriteString(config)

		if err != nil {
			fmt.Println(err)
			return false
		}
		return true

	case module == "rke1":
		config = config + components.ResourceEC2CloudCredentials + components.RKE1Cluster + components.RKE1NodeTemplate
		poolConfig := ``
		num := 1
		for _, pool := range nodePools {
			poolNum := strconv.Itoa(num)
			quantity := strconv.Itoa(pool.Quantity)
			poolConfig = poolConfig + components.RKE1NodePoolPrefix + poolNum + components.RKE1NodePoolSpecs1 + poolNum + components.RKE1NodePoolSpecs2 + quantity + components.RKE1NodePoolSpecs3 + pool.Cp + components.RKE1NodePoolSpecs4 + pool.Etcd + components.RKE1NodePoolSpecs5 + pool.Wkr + components.RKE1NodePoolSuffix
			num = num + 1
		}
		config = config + poolConfig

		_, err = f.WriteString(config)

		if err != nil {
			fmt.Println(err)
			return false
		}
		return true

	case module == "rke2" || module == "k3s":
		config = config + components.DataEC2CloudCredentials + components.MachineConfigV2
		poolConfig := ``
		num := 1
		for _, pool := range nodePools {
			poolNum := strconv.Itoa(num)
			quantity := strconv.Itoa(pool.Quantity)
			poolConfig = poolConfig + components.V2MachinePoolsPrefix + poolNum + components.V2MachinePoolsSpecs1 + pool.Cp + components.V2MachinePoolsSpecs2 + pool.Etcd + components.V2MachinePoolsSpecs3 + pool.Wkr + components.V2MachinePoolsSpecs4 + quantity + components.V2MachinePoolsSuffix
			num = num + 1
		}
		config = config + components.V2ClusterPrefix + poolConfig + components.V2ClusterSuffix

		_, err = f.WriteString(config)

		if err != nil {
			fmt.Println(err)
			return false
		}
		return true

	case module != "aks" || module != "rke1" || module != "rke2" || module != "k3s":
		fmt.Printf("\nModule does not exist; check for possible typo")
		return false

	default:
		return false
	}
}
