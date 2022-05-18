package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetClusterNodeCount(hostURL string, clusterID string, token string) int {

	type ClusterSpecs struct {
		NodeCount int    `json:"nodeCount"`
	}

	url := fmt.Sprintf("%s/v3/clusters/%s", hostURL, clusterID)

	var clusterSpecs ClusterSpecs

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

	jsonErr := json.NewDecoder(res.Body).Decode(&clusterSpecs)
	if err != nil {
		fmt.Printf("%v", jsonErr)
	} 
	clusterSpec := clusterSpecs.NodeCount
	return clusterSpec
}