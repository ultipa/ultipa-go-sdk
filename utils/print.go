package utils

import (
	"encoding/json"
	"log"
)

// PrintJSON prints any interface to json string in cmd
func PrintJSON(obj interface{}) string {
	bytes, err := json.Marshal(obj)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(bytes))

	return string(bytes)
}
