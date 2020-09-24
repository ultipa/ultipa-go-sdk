package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListGraph(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)
	graphs := conn.ListGraph(nil)

	if len(graphs) < 1 {
		log.Fatalln("graph not found")
	}
}

func TestGetGraph(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)
	rs := conn.GetGraph("default", nil)
	log.Println(rs.Graph.Name, rs.Graph.Id, rs.Graph.TotalNodes, rs.Graph.TotalEdges)
}

func TestCreateGraph(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)
	newName := "go_new_graph"
	conn.CreateGraph(newName, nil)
	graphs := conn.ListGraph(nil)
	rs := false
	for _, graph := range graphs {
		// log.Println(graph.Id, graph.Name, graph.TotalEdges, graph.TotalNodes)
		if graph.Name == newName {
			rs = true
		}
	}

	assert.Equal(t, rs, true, "new graph should be found")
}

func TestDropGraph(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)
	newName := "go_new_graph"
	conn.DropGraph(newName, nil)
	graphs := conn.ListGraph(nil)
	rs := true
	for _, graph := range graphs {
		// log.Println(graph.Id, graph.Name, graph.TotalEdges, graph.TotalNodes)
		if graph.Name == newName {
			rs = false
		}
	}

	assert.Equal(t, rs, true, "new graph should be not found")
}
