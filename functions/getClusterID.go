package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetClusterID(hostURL string, clusterName string, token string) string {
	type ClusterSpecs struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	type clusterResponse struct {
		Clusters []ClusterSpecs `json:"data"`
	}

	url := fmt.Sprintf("%s/v3/clusters", hostURL)

	var clusters clusterResponse

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

	jsonErr := json.NewDecoder(res.Body).Decode(&clusters)
	if err != nil {
		fmt.Printf("%v", jsonErr)
	}

	var clusterSpec string

	for i := 0; i < len(clusters.Clusters); i++ {
		if clusters.Clusters[i].Name == clusterName {
			clusterSpec = clusters.Clusters[i].Id
		}
	}

	return clusterSpec
}

// To allow tests to run in parallel, instead of grabbing first cluster object from list of clusters,
// Modify code accept `clusterName` as a parameter and loop through cluster objects and check that
// clusterName matches, if so, return id from that cluster object.  This should always target the
// intended cluster and would allow tests to run in parallel without conflict
//
// Solution has been implemented but not yet tested with multiple clusters in parallel
