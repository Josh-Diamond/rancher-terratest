package functions

import (
	"fmt"
	"os"
)

func UpdateConfigTF(config string, module string) {
	f, err := os.Create("../../modules/" + module + "/main.tf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", f)

	defer f.Close()

	_, err = f.WriteString(config)

	if err != nil {
		fmt.Println(err)
	}
}

// experimental - UpdateConfigTF will take in a string literal of the updated main.tf config and overwrite the previous main.tf
// Thought here was to be able to let terraform add/delete node pools and manipulate resources, instead of building API calls using Go