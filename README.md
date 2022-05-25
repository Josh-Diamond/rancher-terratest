# rancher-terratest

Automated tests for Rancher using Terraform + Terratest


Note: Tests that timeout will likely not have called `terraform destroy`, and as a result, resources will not be cleaned up properly. Extending the test timeout is a best practice. The default timeout when testing with Go is 10 mins; to extend timeout, add `-timeout <int>m` when running tests; e.g. `go test <testfile>.go -timeout 30m`


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
  - parameters - (`url string`, `clusterID string`, `bearer token string`); returns `string`
  - description - returns rancher's server version
- **GetUserID**:
  - parameters - (`url string`, `bearer token string`); returns `string`
  - description - returns admin user id
- **OutputToInt**:
  - parameters - (`output string`); returns `int`
  - description - returns tf output as type int
  - note - tf outputs values as type string; this is not always desired, as actual values coming from rancher server will be various types
- **UpdateConfig**:
  - parameters - (`config string`, `module string`)
  - description - accepts a new "main.tf" config for a specified module and updates the existing main.tf file with new config
  - note - when using with terratest, you will have to run terraform.Appy() afterwards to initiate the updates
- **WaitForActiveCluster**:
  - parameters - (`url string`, `clusterName string`, `bearer token string`)
  - description - waits until cluster is in an active state and ready-to-test before continuing
  - note - required for RKE1; must instantiate in test function after TF `init + apply` and before executing tests
