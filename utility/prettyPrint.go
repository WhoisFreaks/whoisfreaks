package utility

import (
	"encoding/json"
	"fmt"
)

func PrintError(errorInfo interface{}) {
	errorJSON, err := json.MarshalIndent(errorInfo, "", "    ")
	if err != nil {
		fmt.Println("Error while Marshaling the Error Object:", err)
		return
	}
	fmt.Println(string(errorJSON))
}

func PrintInfo(infoObject interface{}) {
	infoJSON, err := json.MarshalIndent(infoObject, "", "    ")
	if err != nil {
		fmt.Println("Error while Marshaling the Object:", err)
		return
	}
	fmt.Println(string(infoJSON))
}
