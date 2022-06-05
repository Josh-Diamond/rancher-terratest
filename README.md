# rancher-terratest

Automated tests for Rancher using Terraform + Terratest

Provisioning:
- AWS Node driver
  - RKE1
  - RKE2
  - K3s
- Hosted
  - AKS


Functions:
- **GetClusterID**: 
  - parameters - (`url string`, `clusterName string`, `bearer token string`); returns `string`
  - description - returns the cluster's id
- **GetClusterName**:
  - parameters - (`url string`, `clusterID string`, `bearer token string`); returns `string`
  - description - returns the cluster's name
- **GetClusterNodeCount**:
  - parameters - (`url string`, `clusterID string`, `bearer token string`); returns `int`
  - description - returns the cluster's node count
- **GetClusterProvider**:
  - parameters - (`url string`, `clusterID string`, `bearer token string`); returns `string`
  - description - returns the cluster's provider
- **GetClusterState**:
  - parameters - (`url string`, `clusterID string`, `bearer token string`); returns `string`
  - description - returns the cluster's current state
- **GetKubernetesVersion**:
  - parameters - (`url string`, `clusterID string`, `bearer token string`); returns `string`
  - description - returns the cluster's kubernetes version
- **GetRancherServerVersion**:
  - parameters - (`url string`, `bearer token string`); returns `string`
  - description - returns rancher's server version
- **GetUserID**:
  - parameters - (`url string`, `bearer token string`); returns `string`
  - description - returns admin user id
- **OutputToInt**:
  - parameters - (`output string`); returns `int`
  - description - returns tf output as type int
  - note - tf outputs values as type string;
- **SetConfigTF**: 
  - parameters - (`module string`, `config []models.Nodepool`; returns `bool`
  - description - sets config of desired module and overwrites exiting main.tf
- **WaitForActiveCluster**:
  - parameters - (`url string`, `clusterName string`, `bearer token string`)
  - description - waits until cluster is in an active state

Testing:
- To add a test with terratest, simple create a new _test.go file in the `tests` folder and begin writing your test!

Note: 
- The default timeout when testing with Go is 10 mins.  
- To extend timeout, add `-timeout <int>m` when running tests; 
  - e.g. `go test <testfile>.go -timeout 45m` || `go test <testfile>.go -timeout 1h`
- Tests that timeout will likely not have resources cleaned up properly. 
- Extending the test timeout is a best practice. 
