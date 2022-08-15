package ibc

import (
	"context"

	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	chantypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
)

func (i IBC) QueryChannelClient(
	ctx context.Context,
	height uint32,
	channelid,
	portid string,
) (
	*clienttypes.IdentifiedClientState,
	error,
) {
	var res clienttypes.IdentifiedClientState
	err := i.client.CallContext(ctx, &res, queryChannelClientMethod, height, channelid, portid)
	if err != nil {
		return &clienttypes.IdentifiedClientState{}, err
	}
	return &res, nil
}

func (i IBC) QueryConnectionChannels(
	ctx context.Context,
	height uint32,
	connectionid string,
) (
	*chantypes.QueryChannelsResponse,
	error,
) {
	var res *chantypes.QueryChannelsResponse
	err := i.client.CallContext(ctx, &res, queryConnectionChannelsMethod, height, connectionid)
	if err != nil {
		return &chantypes.QueryChannelsResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryChannels(ctx context.Context) (
	*chantypes.QueryChannelsResponse,
	error,
) {
	var res *chantypes.QueryChannelsResponse
	err := i.client.CallContext(ctx, &res, queryChannelsMethod)
	if err != nil {
		return &chantypes.QueryChannelsResponse{}, err
	}
	return res, nil
}
