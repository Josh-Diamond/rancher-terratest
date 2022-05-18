
# Rancher specific variable section.
variable rancher_api_url {}
variable rancher_admin_bearer_token {}
variable cloud_credential_name {}

# AWS specific variables.
variable aws_access_key {}
variable aws_secret_key {}
variable aws_ami {}
variable aws_region {}
variable aws_security_group_name {}
variable aws_subnet_id {}
variable aws_vpc_id {}
variable aws_zone_letter {}

# RKE2/k3s specific variables.
variable machine_config_name {}
variable cluster_name {}
variable enable_network_policy {}
variable default_cluster_role_for_project_members {}
variable kubernetes_version {}