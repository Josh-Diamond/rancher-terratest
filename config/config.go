package config

import (
	"github.com/josh-diamond/rancher-terratest/models"
)

// This is your workspace, feel free to build out multiple configurations by declaring
// K8sVersions and desired NodePools, then apply them whenever you'd like in your test
// Tests may be added in the tests package with file name suffix _test.go 
// Failure to do so and your tests will not be recognized as such
// [ e.g. rancher > terratest > tests > module > testfile_test.go ]

// Modules
var Aks  = "aks"
var K3s  = "k3s"
var Rke1 = "rke1"
var Rke2 = "rke2"

// K8s versions
var AKSK8sVersion1235   = "1.23.5"
var AKSK8sVersion1226   = "1.22.6"
var AKSK8sVersion1219   = "1.21.9"

var K3sK8sVersion1236   = "v1.23.6+k3s1"
var K3sK8sVersion1229   = "v1.22.9+k3s1"
var K3sK8sVersion12112  = "v1.21.12+k3s1"

var RKE1K8sVersion1236  = "v1.23.6-rancher1-1"
var RKE1K8sVersion1229  = "v1.22.9-rancher1-1"
var RKE1K8sVersion12112 = "v1.21.12-rancher1-1"

var RKE2K8sVersion1236  = "v1.23.6+rke2r2"
var RKE2K8sVersion1229  = "v1.22.9+rke2r2"
var RKE2K8sVersion12112 = "v1.21.12+rke2r2"


// Customize your desired node pools using unique names
// *examples_here*
// Declare new variable for nodepools 
// e.g. 'var NewNodePools []models.Nodepool'
// Create a function that appends the new desired node pools to the new declared slice NewNodePools
//
// e.g. func BuildNewNodePools() {
//           NewNodePools = append(NewNodePools, unique_pool1, unique_pool2, unique_pool3)
//      }
//
// NewNodePools variable and function that appends to NewNodePools must beginning with a capital letter
// Variables and functions failing to meet this requirement will not export and will be unaccessible during tests
// Repeat steps for multiple configs to test against

// Nodes3_Etcd1_Cp1_Wkr1
var etcd_1 = models.Nodepool{
	Quantity: 1,
	Etcd:     "true",
	Cp:       "false",
	Wkr:      "false",
}

var cp_1 = models.Nodepool{
	Quantity: 1,
	Etcd:     "false",
	Cp:       "true",
	Wkr:      "false",
}

var wkr_1 = models.Nodepool{
	Quantity: 1,
	Etcd:     "false",
	Cp:       "false",
	Wkr:      "true",
}

var Nodes3_Etcd1_Cp1_Wkr1 []models.Nodepool

func Build_Nodes3_Etcd1_Cp1_Wkr1() {
	Nodes3_Etcd1_Cp1_Wkr1 = append(Nodes3_Etcd1_Cp1_Wkr1, etcd_1, cp_1, wkr_1)
}

// Nodes8_HACluster
var etcd_3 = models.Nodepool{
	Quantity: 3,
	Etcd:     "true",
	Cp:       "false",
	Wkr:      "false",
}

var cp_2 = models.Nodepool{
	Quantity: 2,
	Etcd:     "false",
	Cp:       "true",
	Wkr:      "false",
}

var wkr_3 = models.Nodepool{
	Quantity: 3,
	Etcd:     "false",
	Cp:       "false",
	Wkr:      "true",
}

var Nodes8_HACluster []models.Nodepool

func Build_Nodes8_HACluster() {
	Nodes8_HACluster = append(Nodes8_HACluster, etcd_3, cp_2, wkr_3)
}

// Nodes6_Etcd3_Cp2_Wkr1
var Nodes6_Etcd3_Cp2_Wkr1 []models.Nodepool

func Build_Nodes6_Etcd3_Cp2_Wkr1() {
	Nodes6_Etcd3_Cp2_Wkr1 = append(Nodes6_Etcd3_Cp2_Wkr1, etcd_3, cp_2, wkr_1)
}
