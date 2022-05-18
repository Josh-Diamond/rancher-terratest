# rancher-terratest

Automated tests for Rancher using Terraform + Terratest

Provisioning:
- AKS Clusters
- RKE1 Clusters
- RKE2 Clusters

Management:
- Coming soon - [Next up: add/delete node pools]
- Idea: create "management" functions, which can create, manipulate, and destroy cluster level resources through API calls with go to be used in tests after inital infrastructer is provisioned

Functions:
- Add_node_pool - [coming soon]
- Delete_node_pool - [coming soon]
- Scale_up_existing_pool - [coming soon]
- Scale_down_existing_pool - [coming soon]
