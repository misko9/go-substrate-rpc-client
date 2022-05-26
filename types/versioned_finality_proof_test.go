package types_test

import (
	"bytes"
	"testing"

	gsrpc "github.com/ComposableFi/go-substrate-rpc-client/v4"
	"github.com/ComposableFi/go-substrate-rpc-client/v4/types"
	"github.com/stretchr/testify/require"
)

func TestVersionedFinalityProof_Decoding(t *testing.T) {
	rococo, err := gsrpc.NewSubstrateAPI("wss://rococo-rpc.polkadot.io")
	require.NoError(t, err)

	hash, err := rococo.RPC.Beefy.GetFinalizedHead()
	require.NoError(t, err)

	block, err := rococo.RPC.Chain.GetBlock(hash)
	require.NoError(t, err)

	for _, v := range block.Justifications {
		if bytes.Equal(v.ConsensusEngineID[:], []byte("BEEF")) {
			versionedFinalityProof := &types.VersionedFinalityProof{}

			err = types.DecodeFromBytes(v.EncodedJustification, versionedFinalityProof)
			require.NoError(t, err)
			t.Logf("%v", versionedFinalityProof.AsCompactSignedCommitment.Unpack())
		}
	}

}
