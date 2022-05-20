package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetClusterProvider(hostURL string, clusterID string, token string) string {

	type clusterSpec struct {
		Provider string `json:"provider"`
	}

	url := fmt.Sprintf("%s/v3/clusters/%s", hostURL, clusterID)

	var spec clusterSpec

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

	jsonErr := json.NewDecoder(res.Body).Decode(&spec)
	if err != nil {
		fmt.Printf("%v", jsonErr)
	}
	provider := spec.Provider
	return provider
}
