package test

import (
	"testing"
)

func TestListGraph(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)
	conn.ListGraph(nil)
}
