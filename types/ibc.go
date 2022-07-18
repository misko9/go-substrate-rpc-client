package types

import (
	"github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

type Proof struct {
	Proof  []byte
	Height types.Height
}

type ConnHandshakeProof struct {
	// Protobuf encoded client state
	ClientState clienttypes.IdentifiedClientState
	// Trie proof for connection state, client state and consensus state
	Proof []byte
	// Proof height
	Height types.Height
}
