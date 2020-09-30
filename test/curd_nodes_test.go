package test

import "testing"

func TestInsertNodes(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)
	nodes := []*map[string]interface{}{
		{
			"name": "test1",
			"age":  2,
		},
		{
			"name": "test2",
			"age":  3,
		},
	}

	conn.InsertNodes(nodes, false, nil)
}

func TestDeleteNodes(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)
	conn.DeleteNodes(map[string]interface{}{
		"name": "test2",
	}, nil)

}
func TestUpdateNodes(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)
	conn.UpdateNodes(map[string]interface{}{
		"name": "test2",
	}, map[string]interface{}{
		"name": "test2go",
	}, nil)

}
