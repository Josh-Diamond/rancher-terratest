package functions

import (
	"fmt"
	"os"
)

// this page will likely be deleted; pending SetConfigTF() completion
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
