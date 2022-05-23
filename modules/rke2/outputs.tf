output "cluster_name" {
  value = rancher2_cluster_v2.rancher2_cluster_v2.name
}

output "host_url" {
  value = var.rancher_api_url
  sensitive = true
}

output "token" {
  value = var.rancher_admin_bearer_token
  sensitive = true
}

output "token_prefix" {
  value = var.token_prefix
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

output "expected_kubernetes_version" {
  value = var.expected_kubernetes_version
}

output "expected_rancher_server_version" {
  value = var.expected_rancher_server_version
}