package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetClusterID(hostURL string, token string) string {
	type clusterSpec struct {
		Id string `json:"id"`
	}
	type clusterResponse struct {
		Clusters []clusterSpec `json:"data"`
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
	id := clusters.Clusters[0].Id
	return id
}

// To allow tests to run in parallel, instead of grabbing first cluster object from list of clusters,
// Modify code accept `clusterName` as a parameter and loop through cluster objects and check that 
// clusterName matches, if so, return id from that cluster object.  This should always target the 
// intended cluster and would allow tests to run in parallel without conflict