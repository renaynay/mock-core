package mock

import "testing"

func TestCreateMockNode(t *testing.T) {
	node := CreateMockNode()
	t.Log("node created: ", node)
}
