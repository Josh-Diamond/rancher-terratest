package functions

import "time"

func WaitForActiveCLuster(hostURL string, clusterName string, token string) {
	id := GetClusterID(hostURL, clusterName, token)
	state := GetClusterState(hostURL, id, token)
	updating := false

	time.Sleep(11 * time.Second)

	for state != "active" {
		for state != "active" && !updating {
			state = GetClusterState(hostURL, id, token)
			time.Sleep(10 * time.Second)
			if state == "updating" {
				updating = true
			}
		}
		state = GetClusterState(hostURL, id, token)
		time.Sleep(10 * time.Second)
	}
}
