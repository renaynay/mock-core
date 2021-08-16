package mock

import (
	celrand "github.com/celestiaorg/celestia-core/libs/rand"
	"github.com/celestiaorg/celestia-core/types"
	"time"
)

type Block struct {
	types.Block
}

func GenerateBlock(peer *Peer, prevBlock Block) Block {
	header := types.Header{
		ChainID: peer.ChainID,
		Height: peer.Height,
		Time: time.Now(),
		LastBlockID: types.BlockID{Hash: prevBlock.Hash()},
		DataHash: celrand.Bytes(100), // TODO how many bytes should this be?

	}

	data := types.Data{

	}

	dah := types.DataAvailabilityHeader{

	}

	commit := new(types.Commit)

	block := Block{
		types.Block{
			Header: header,
			Data: data,
			DataAvailabilityHeader: dah,
			LastCommit: commit,
		},
	}
	return block
}

// func GenerateBlock
// func GenerateHeader

