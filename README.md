# rancher-terratest

Automated tests for Rancher using Terraform + Terratest

Provisioning:
- AKS Clusters
- RKE1 Clusters
- RKE2 Clusters

Management:
- Coming soon 
- Idea: Define "management" functions, which can create, manipulate, and destroy cluster level resources through API calls with go to be used in tests after inital infrastructer is provisioned

Functions:
- Get_cluster_specs - [grabs cluster name, nodeCount, provider, and state] 
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
