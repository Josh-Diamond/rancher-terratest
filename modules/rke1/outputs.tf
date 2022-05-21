output "cluster_name_rke1" {
  value = rancher2_cluster.rancher2_cluster.name
}

output "host_url" {
  value = var.rancher_api_url
  sensitive = true
}

output "bearer_token" {
  value = var.rancher_admin_bearer_token
  sensitive = true
}