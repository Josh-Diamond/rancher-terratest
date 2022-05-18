package test

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetClusterSpecs(hostURL string, clusterID string, token string) {

	type ClusterSpecs struct {
		Name      string `json:"name"`
		NodeCount int    `json:"nodeCount"`
		Provider  string `json:"provider"`
		State     string `json:"state"`
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
	} else {
		fmt.Printf("Name: %v\n", clusterSpecs.Name)
		fmt.Printf("NodeCount: %v\n", clusterSpecs.NodeCount)
		fmt.Printf("Provider: %v\n", clusterSpecs.Provider)
		fmt.Printf("State: %v\n", clusterSpecs.State)
		// fmt.Printf("%v",clusterSpecs)
	}
}
