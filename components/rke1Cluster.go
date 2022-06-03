package components

var RKE1Cluster = `resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  rke_config {
    kubernetes_version = var.kubernetes_version
	  network {
	    plugin = var.network_plugin
	  }
  }
}

`