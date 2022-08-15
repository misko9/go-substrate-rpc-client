package ibc

import (
	"context"

	"github.com/ComposableFi/go-substrate-rpc-client/v4/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

func (i IBC) QueryClientStateResponse(
	ctx context.Context,
	height int64,
	srcClientID string,
) (
	*clienttypes.QueryClientStateResponse,
	error,
) {
	var res *clienttypes.QueryClientStateResponse
	err := i.client.CallContext(ctx, &res, queryClientStateMethod, height, srcClientID)
	if err != nil {
		return &clienttypes.QueryClientStateResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryClientConsensusState(
	ctx context.Context,
	clientid string,
	revisionHeight,
	revisionNumber uint64,
	latestConsensusState bool) (
	*clienttypes.QueryConsensusStateResponse,
	error,
) {
	var res *clienttypes.QueryConsensusStateResponse
	err := i.client.CallContext(ctx, &res,
		queryClientConsensusStateMethod,
		clientid,
		revisionHeight,
		revisionNumber,
		latestConsensusState)
	if err != nil {
		return &clienttypes.QueryConsensusStateResponse{}, err
	}
	return res, nil
}
func (i IBC) QueryUpgradedClient(
	ctx context.Context,
	height int64,
) (*clienttypes.QueryClientStateResponse, error) {
	var res *clienttypes.QueryClientStateResponse
	err := i.client.CallContext(ctx, &res, queryUpgradedClientMethod, height)
	if err != nil {
		return &clienttypes.QueryClientStateResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryUpgradedConsState(
	ctx context.Context,
	height int64,
) (
	*clienttypes.QueryConsensusStateResponse,
	error,
) {
	var res *clienttypes.QueryConsensusStateResponse
	err := i.client.CallContext(ctx, &res, queryUpgradedConnectionStateMethod, height)
	if err != nil {
		return &clienttypes.QueryConsensusStateResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryClients(ctx context.Context) (
	clienttypes.IdentifiedClientStates,
	error,
) {
	var res clienttypes.IdentifiedClientStates
	var x interface{}
	err := i.client.CallContext(ctx, &x, queryClientsMethod)
	if err != nil {
		return clienttypes.IdentifiedClientStates{}, err
	}
	return res, nil
}

func (i IBC) QueryNewlyCreatedClients(
	ctx context.Context,
	blockHash types.H256,
) (
	[]clienttypes.IdentifiedClientState,
	error,
) {
	var res []clienttypes.IdentifiedClientState
	err := i.client.CallContext(ctx, &res, queryNewlyCreatedClientsMethod)
	if err != nil {
		return []clienttypes.IdentifiedClientState{}, err
	}
	return res, nil
}
