package test

import (
	"log"
	"testing"
)

func TestShowJob(t *testing.T) {
	jobs, err := client.ShowJob("", nil)
	//jobs, err := client.Uql("db.backup.show()", nil)
	if err != nil {
		t.Fatalf("show jobs error %v", err)
	}

	//t.Log(http.GetJobResponseFromUqlResponse(jobs))
	for _, job := range jobs {
		log.Println(job)
	}
}
