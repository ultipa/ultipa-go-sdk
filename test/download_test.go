package test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	receive := func(data []byte) error {
		fmt.Printf("%s\n", data)
		return nil
	}
	fileName := "aggregations"
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	err = client.DownloadAlgoResultFile(fileName, "4", nil, receive)
	log.Println(err.Error())
}
