# Rancher variables
variable rancher_api_url {}
variable rancher_admin_bearer_token {}

# Azure credential variables
variable cloud_credential_name {}
variable azure_client_id {}
variable azure_client_secret {}
variable azure_subscription_id {}

# AKS variables
variable cluster_name {}
variable resource_group {}
variable resource_location {}
variable dns_prefix {}
variable kubernetes_version {}
variable network_plugin {}
variable availability_zones {}
variable orchestrator_version {}
variable os_disk_size_gb {}
variable vm_size {}

# Testing variables
variable token_prefix {}
variable expected_node_count {}
variable expected_provider {}
variable expected_state {}
variable expected_kubernetes_version {}
variable expected_rancher_server_version {}