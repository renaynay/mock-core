package mock

import (
	"github.com/celestiaorg/celestia-core/types"
	"github.com/renaynay/celestia-core/testutils"
	"math/rand"
	"time"
)

var (
	maxTxsInBlock = 100
	maxMessagesInBlock = 20

	maxTxOrMessageSize = 100

	maxISRAmount = 10
	maxEvidenceAmount = 5
)

type Block struct {
	types.Block
}

func GenerateGenesisBlock() (Block, error) {
	// TODO
	return Block{}, nil
}

// GenerateBlock generates a full Celestia dummy block.
func GenerateBlock(peer *Peer, prevBlock Block) (Block, error) {
	var block Block
	// generate block data
	data, err := testutils.GenerateRandomBlockData(rand.Intn(maxTxsInBlock), rand.Intn(maxISRAmount),
		rand.Intn(maxEvidenceAmount), rand.Intn(maxMessagesInBlock), maxTxOrMessageSize)
	if err != nil {
		return Block{}, err
	}
	block.Data = data
	// generate header
	header := types.Header{
		ChainID: peer.ChainID,
		Height: prevBlock.Height + 1,
		Time: time.Now(),
		LastBlockID: prevBlock.LastBlockID,
		// TODO LastCommitHash?
		// TODO DataHash
		// TODO NumOriginalDataShares
		// TODO ValidatorsHash
		// TODO NextValidatorsHash
		// TODO ConsensusHash
		// TODO AppHash
		// TODO LastResultsHash
		EvidenceHash: data.Evidence.Hash(),
		// TODO ProposerAddress
	}
	block.Header = header
	// generate DAH
	// TODO
	// generate last commit?
	// TODO
	return block, nil
}

func (b *Block) GetHeader() types.Header {

}

// func GenerateHeader


