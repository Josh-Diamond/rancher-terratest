package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProvisioningRke2ClusterID(hostURL string, clusterName string, token string) string {
	type clusterSpecs struct {
		ClusterName string `json:"clusterName"`
	}
	type statusResponse struct {
		Status clusterSpecs `json:"status"`
	}

	url := fmt.Sprintf("%s/v1/provisioning.cattle.io.clusters/fleet-default/%s", hostURL, clusterName)

	var status statusResponse

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("%v", err)
	}

	req.Header = http.Header{
		"Authorization": []string{token},
	}

	res, clientErr := client.Do(req)
	if clientErr != nil {
		fmt.Printf("%v", clientErr)
	}

	defer res.Body.Close()

	jsonErr := json.NewDecoder(res.Body).Decode(&status)
	if err != nil {
		fmt.Printf("%v", jsonErr)
	}
	clusterSpec := status.Status.ClusterName
	return clusterSpec
}

// Not available for RKE1; needs to be tested with AKS before determining this is RKE2/K3s specific