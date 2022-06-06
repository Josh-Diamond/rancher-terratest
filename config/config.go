package config

import (
	"github.com/josh-diamond/rancher-terratest/models"
)

// This is your workspace, feel free to build out multiple configurations with
// The instruction provided below, then apply them whenever you'd like in your terratest
// Custom terratests may be added as well, in the tests package with file name suffix _test.go
// Failure to do so and your tests will not be recognized as such

// Modules
var Aks  = "aks"
var Rke1 = "rke1"
var Rke2 = "rke2"
var K3s  = "k3s"

// K8s versions
var AKSK8sVersion1219 = "1.21.9"
var AKSK8sVersion1226 = "1.22.6"

var RKE1K8sVersion1229 = "v1.22.9-rancher1-1"
var RKE1K8sVersion1236 = "v1.23.6-rancher1-1"

var RKE2K8sVersion1229 = "v1.22.9+rke2r2"
var RKE2K8sVersion1236 = "v1.23.6+rke2r2"

var K3sK8sVersion1229 = "v1.22.9+k3s1"
var K3sK8sVersion1236 = "v1.23.6+k3s1"

// Customize your desired node pools for Config1
// Update append() on (line 40) to include desired node pools
// For multiple configurations, repeat steps for desired node pools
// Use unique names, then create a variable for the new config of type []models.Nodepool
// e.g. var Config2 []models.Nodepool
// Next, create a function, e.g. BuildConfig2() that appends
// the new desired node pools to the new config
// New config variable must beginning with a capital letter
// It will not export otherwise and will be unaccessible during the terratest
// Repeat steps for multiple configs to test against

// NodePools1
var pool1 = models.Nodepool{
	Quantity: 1,
	Etcd:     "true",
	Cp:       "false",
	Wkr:      "false",
}

var pool2 = models.Nodepool{
	Quantity: 1,
	Etcd:     "false",
	Cp:       "true",
	Wkr:      "false",
}

var pool3 = models.Nodepool{
	Quantity: 1,
	Etcd:     "false",
	Cp:       "false",
	Wkr:      "true",
}

var NodePools1 []models.Nodepool

func BuildNodePools1() {
	NodePools1 = append(NodePools1, pool1, pool2, pool3)
}

// NodePools2
var pool4 = models.Nodepool{
	Quantity: 3,
	Etcd:     "true",
	Cp:       "false",
	Wkr:      "false",
}

var pool5 = models.Nodepool{
	Quantity: 2,
	Etcd:     "false",
	Cp:       "true",
	Wkr:      "false",
}

var pool6 = models.Nodepool{
	Quantity: 3,
	Etcd:     "false",
	Cp:       "false",
	Wkr:      "true",
}

var NodePools2 []models.Nodepool

func BuildNodePools2() {
	NodePools2 = append(NodePools2, pool4, pool5, pool6)
}

// NodePools3
var pool7 = models.Nodepool{
	Quantity: 3,
	Etcd:     "true",
	Cp:       "false",
	Wkr:      "false",
}

var pool8 = models.Nodepool{
	Quantity: 2,
	Etcd:     "false",
	Cp:       "true",
	Wkr:      "false",
}

var pool9 = models.Nodepool{
	Quantity: 1,
	Etcd:     "false",
	Cp:       "false",
	Wkr:      "true",
}

var NodePools3 []models.Nodepool

func BuildNodePools3() {
	NodePools3 = append(NodePools3, pool7, pool8, pool9)
}
