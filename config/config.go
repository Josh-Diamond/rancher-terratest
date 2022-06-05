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

// Config1
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

var Config1 []models.Nodepool

func BuildConfig1() {
	Config1 = append(Config1, pool1, pool2, pool3)
}

// Config2
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

var Config2 []models.Nodepool

func BuildConfig2() {
	Config2 = append(Config2, pool4, pool5, pool6)
}

// Config3
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

var Config3 []models.Nodepool

func BuildConfig3() {
	Config3 = append(Config3, pool7, pool8, pool9)
}
