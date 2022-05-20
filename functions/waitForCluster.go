package functions

import "time"

func WaitForCLuster(hostURL string, token string) {
	id := GetClusterID(hostURL, token)
	state := GetClusterState(hostURL, id, token)
	updating := false

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
