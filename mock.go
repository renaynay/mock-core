package mock

type MockNode struct {
	Peer *Peer
	PeerAddr string // dial address of the celestia node "peer"
	Server *Server
}

func CreateMockNode(endpoint string) (*MockNode, error) {
	node := new(MockNode)
	node.Server = NewServer(endpoint)
	return node, nil
}

func (m *MockNode) Start() error {
	return m.Server.Start()
}

func (m *MockNode) Stop() error {
	return m.Server.Stop()
}

func (m *MockNode) EmitBlock() error {
}

// func EmitBlock
// func EmitHeader
