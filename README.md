# rancher-terratest

Automated tests for Rancher using Terraform + Terratest

Provisioning:
- AKS
- RKE1
- RKE2
- K3s
- EKS - [W.I.P]
- GKE - [comming soon]



Functions:
- **GetClusterID**: 
  - parameters (`url string`, `bearer token string`); returns `string`
- **GetClusterName**:
  - parameters (`url string`, `clusterID string`, `bearer token string`); returns `string`
- **GetClusterNodeCount**:
  - parameters (`url string`, `clusterID string`, `bearer token string`); returns `int`
- **GetClusterProvider**:
  - parameters (`url string`, `clusterID string`, `bearer token string`); returns `string`
- **GetClusterState**:
  - parameters (`url string`, `clusterID string`, `bearer token string`); returns `string`
- **GetProvisioningRke2ClusterID**:
  - parameters (`url string`, `clusterName string`, `bearer token string`); returns `string`
- **WaitUntilRke2ClustersActive**:
  - parameters (`url string`, `clusterName string`, `bearer token string`); returns `string`

- Add_node_pool - [coming soon]
- Delete_node_pool - [coming soon]
- Scale_up_existing_pool - [coming soon]
- Scale_down_existing_pool - [coming soon]
- Take_etcd_snapshot - [coming soon]
- Restore_etcd_snapshot - [coming soon]
- Deploy_workload - [coming soon]
- Delete_workload - [coming soon]
- Create_project - [coming soon]
- Delete_project - [coming soon]
- Create_namespace - [coming soon]
- Delete_namespace - [coming soon]
- Create_user - [coming soon]
- Delete_user - [coming soon]
- Edit_cluster_config - [coming soon]- patch any data passed into function 
- Install_chart - [coming soon]
