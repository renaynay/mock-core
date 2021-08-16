package mock

import (
	"net"
	"testing"
)

func TestCreateAndStartMockNode(t *testing.T) {
	listenAddr := "0.0.0.0:7878"
	node, err := CreateMockNode(listenAddr)
	if err != nil {
		t.Fatal(err)
	}
	if err := node.Server.Start(); err != nil {
		t.Fatal(err)
	}
	defer node.Server.Stop()
	t.Log("node created: ", node.Server.Listener.Addr())

	conn, err := net.Dial("tcp", node.Server.Listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dial successful")
	defer conn.Close()
}