terraform {
  required_providers {
    rancher2 = {
	  source  = "rancher/rancher2"
	  version = "1.22.2"
	 }
  }
}

provider "rancher2" {
  api_url   = var.rancher_api_url
  token_key = var.rancher_admin_bearer_token
}

resource "rancher2_cloud_credential" "rancher2_cloud_credential" {
  name = var.cloud_credential_name
  amazonec2_credential_config {
    access_key = var.aws_access_key
    secret_key = var.aws_secret_key
    default_region = var.aws_region
  }
}

resource "rancher2_cluster" "rancher2_cluster" {
  name = var.cluster_name
  rke_config {
	  kubernetes_version = var.kubernetes_version
	  network {
	    plugin = var.network_plugin
	  }
  }
}
	  
resource "rancher2_node_template" "rancher2_node_template" {
  name = var.node_template_name
  amazonec2_config {
    access_key     = var.aws_access_key
	  secret_key     = var.aws_secret_key
	  ami            = var.aws_ami_w_docker
	  region         = var.aws_region
	  security_group = [var.aws_security_group_name]
	  subnet_id      = var.aws_subnet_id
	  vpc_id         = var.aws_vpc_id
	  zone           = var.aws_zone_letter
	  root_size      = var.aws_root_size
	  instance_type  = var.aws_instance_type
  }
}

resource "rancher2_node_pool" "pool1" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool1"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 1
  control_plane    = true
  etcd             = true 
  worker           = true 
}

resource "rancher2_node_pool" "pool2" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool2"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 2
  control_plane    = true
  etcd             = false 
  worker           = true 
}

resource "rancher2_node_pool" "pool3" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool3"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 3
  control_plane    = false
  etcd             = true 
  worker           = true 
}

resource "rancher2_node_pool" "pool4" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool4"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 4
  control_plane    = true
  etcd             = true 
  worker           = false 
}

resource "rancher2_node_pool" "pool5" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool5"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 5
  control_plane    = false
  etcd             = false 
  worker           = true 
}

resource "rancher2_node_pool" "pool6" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool6"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 6
  control_plane    = true
  etcd             = false 
  worker           = false 
}

resource "rancher2_node_pool" "pool7" {
  cluster_id       = rancher2_cluster.rancher2_cluster.id
  name             = "pool7"
  hostname_prefix  = var.node_hostname_prefix
  node_template_id = rancher2_node_template.rancher2_node_template.id
  quantity         = 7
  control_plane    = true
  etcd             = true 
  worker           = true 
}