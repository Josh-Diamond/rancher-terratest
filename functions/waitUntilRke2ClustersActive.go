package functions


func WaitUntilRke2ClustersActive(hostURL string, clusterName string, token string) bool {
	id := GetProvisioningRke2ClusterID(hostURL, clusterName, token)
	state := GetClusterState(hostURL, id, token)
	updating := false

	for state != "active" {
		for state != "active" && !updating {
			state = GetClusterState(hostURL, id, token)
			if state == "updating" {
				updating = true
			}
		}
		state = GetClusterState(hostURL, id, token)
	}

	return true

}