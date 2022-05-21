output "cluster_name_rke2" {
  value = rancher2_cluster_v2.rancher2_cluster_v2.name
}

output "host_url" {
  value = var.rancher_api_url
  sensitive = true
}

output "bearer_token" {
  value = var.rancher_admin_bearer_token
  sensitive = true
}