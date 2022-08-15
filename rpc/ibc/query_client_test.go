package ibc_test

import (
	"context"
	"testing"

	client "github.com/ComposableFi/go-substrate-rpc-client/v4"
	"github.com/ComposableFi/go-substrate-rpc-client/v4/config"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/stretchr/testify/require"
)

func TestQueryClient(t *testing.T) {
	cl, err := client.NewSubstrateAPI(config.Default().RPCURL)
	require.NoError(t, err)

	clients, err := cl.RPC.IBC.QueryClients(context.Background())
	require.NoError(t, err)
	require.Equal(t, clients, clienttypes.IdentifiedClientStates{})
}
