package sdk

/*
 * Download files from ultipa-server
 */
import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"

	// "strings"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

func Download(client Client, path string, outPath string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24*7)
	defer cancel()

	msg, err := client.DownloadFile(ctx, &ultipa.DownloadFileRequest{
		FilePathName: path,
	})

	if err != nil {
		log.Printf("[Error] download file error: %v", err)
	}

	merr := os.MkdirAll("./downloads", os.ModePerm)
	if merr != nil {
		fmt.Printf("Create path Failed: %v ", merr)
		return
	}

	// var blob []byte
	fmt.Println("download to " + outPath)
	bar := pb.Full.Start(0)
	bar.Set(pb.Bytes, true)

	for {
		c, err := msg.Recv()

		if err != nil {
			if err == io.EOF {
				// log.Printf("File Download OK %v", len(blob))
				break
			} else {
				log.Printf("Failed %v", err)
				break
			}

			// panic(err)
		}
		bar.SetTotal(int64(c.TotalSize))
		bar.SetCurrent(int64(len(c.Chunk)) + bar.Current())
		ioutil.WriteFile(outPath, c.Chunk, 0644)
	}

	bar.Finish()

}
