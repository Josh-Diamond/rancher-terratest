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
