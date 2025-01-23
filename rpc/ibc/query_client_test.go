package ibc_test

import (
	"context"
	"testing"

	clienttypes "github.com/cosmos/ibc-go/v9/modules/core/02-client/types"
	client "github.com/misko9/go-substrate-rpc-client/v4"
	"github.com/misko9/go-substrate-rpc-client/v4/config"
	"github.com/stretchr/testify/require"
)

func TestQueryClient(t *testing.T) {
	cl, err := client.NewSubstrateAPI(config.Default().RPCURL)
	require.NoError(t, err)

	clients, err := cl.RPC.IBC.QueryClients(context.Background())
	require.NoError(t, err)
	require.Equal(t, clients, clienttypes.IdentifiedClientStates{})
}
