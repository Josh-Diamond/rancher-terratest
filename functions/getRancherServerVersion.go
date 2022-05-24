package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// From local cluster
func GetRancherServerVersion(hostURL string, token string) string {
	type nestedData struct {
		Version string `json:"seen-whatsnew"`
	}

	type clusterSpecs struct {
		Id  string `json:"id"`
		Dataa nestedData `json:"data"`
	}

	type clusterResponse struct {
		Data []clusterSpecs `json:"data"`
	}

	url := fmt.Sprintf("%s/v1/userpreferences", hostURL)

	var response clusterResponse

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

	jsonErr := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Printf("%v", jsonErr)
	}

	clippedID := strings.Trim(response.Data[0].Dataa.Version, `"`)

	return clippedID
}




// local seems to be more reliable as it is available as soon as the server is launched; 
// Downstream needs cluster before it can retrieved
//
//
// From downstream cluster
//
// func GetRancherServerVersion(hostURL string, clusterID string, token string) string {
// 	type clusterSpecs struct {
// 		Name  string `json:"name"`
// 		Value string `json:"value"`
// 	}
// 	type clusterResponse struct {
// 		AgentEnvVars []clusterSpecs `json:"appliedAgentEnvVars"`
// 	}

// 	url := fmt.Sprintf("%s/v3/clusters/%s", hostURL, clusterID)

// 	var agentEnvVars clusterResponse

// 	client := http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	req.Header = http.Header{
// 		"Authorization": []string{token},
// 	}

// 	res, clientErr := client.Do(req)
// 	if clientErr != nil {
// 		fmt.Printf("%v", clientErr)
// 	}

// 	defer res.Body.Close()

// 	jsonErr := json.NewDecoder(res.Body).Decode(&agentEnvVars)
// 	if err != nil {
// 		fmt.Printf("%v", jsonErr)
// 	}

// 	var serverVersion string

// 	for _, vars := range agentEnvVars.AgentEnvVars {
// 		if vars.Name == "CATTLE_SERVER_VERSION" {
// 			serverVersion = vars.Value
// 		}
// 	}

// 	return serverVersion
// }
