package test

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	fileName := "data/ids"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	receive := func(data []byte) error {
		_, err = file.Write(data)

		if err != nil {
			return err
		}
		return nil
	}
	err = client.DownloadAlgoResultFile(fileName, "1", nil, receive)
	if err != nil {
		t.Error(err)
	}
}

func TestDownloadAll(t *testing.T) {
	receive := func(data []byte, fileName string) error {

		file, err := os.OpenFile("./data/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		_, err = file.Write(data)

		if err != nil {
			return err
		}

		return nil
	}
	err := client.DownloadAllAlgoResultFile("1", nil, receive)
	if err != nil {
		t.Error(err)
	}
}
