output "cluster_name_aks" {
  value = rancher2_cluster.rancher2_cluster.name
}

output "host_url" {
  value = var.rancher_api_url
  sensitive = true
}

output "token" {
  value = var.rancher_admin_bearer_token
  sensitive = true
}

output "token_type" {
  value = var.token_type
  sensitive = true
}

output "expected_node_count" {
  value = var.expected_node_count
}

output "expected_provider" {
  value = var.expected_provider
}

output "expected_state" {
  value = var.expected_state
}