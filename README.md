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
  - parameters - (`url string`, `bearer token string`); returns `string`
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
- **WaitForCluster**:
  - parameters - (`url string`, `clusterName string`, `bearer token string`); returns `string`
  - description - will wait until cluster is in an active state and ready-to-test before continuing

- Generate_token - [coming soon] 
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
