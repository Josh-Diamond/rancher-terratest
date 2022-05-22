package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRancherServerVersion(hostURL string, clusterID string, token string) string {
	type clusterSpecs struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	type clusterResponse struct {
		AgentEnvVars []clusterSpecs `json:"appliedAgentEnvVars"`
	}

	url := fmt.Sprintf("%s/v3/clusters/%s", hostURL, clusterID)

	var variables clusterResponse

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

	jsonErr := json.NewDecoder(res.Body).Decode(&variables)
	if err != nil {
		fmt.Printf("%v", jsonErr)
	}

	var serverVersion string

	for _, vars := range variables.AgentEnvVars {
		if vars.Name == "CATTLE_SERVER_VERSION" {
			serverVersion = vars.Value
		}
	}

	return serverVersion
}
// Needs to be tested