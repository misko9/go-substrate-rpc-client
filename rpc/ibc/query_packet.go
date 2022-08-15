package ibc

import (
	"context"

	chantypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
)

func (i IBC) QueryAcknowledgements(
	ctx context.Context,
	height uint64,
	channelid,
	portid string) (
	[][]byte,
	error,
) {
	var res [][]byte
	err := i.client.CallContext(ctx, &res, queryAcknowledgementsMethod, height, channelid, portid)
	if err != nil {
		return [][]byte{}, err
	}
	return res, nil
}

func (i IBC) QueryPackets(
	ctx context.Context,
	channelid,
	portid string,
	seqs []uint64,
) (
	[]chantypes.Packet,
	error,
) {
	var res []chantypes.Packet
	err := i.client.CallContext(ctx, &res, queryPacketsMethod, channelid, portid, seqs)
	if err != nil {
		return []chantypes.Packet{}, err
	}
	return res, nil
}

func (i IBC) QueryPacketCommitments(
	ctx context.Context,
	height uint64,
	channelid,
	portid string) (
	*chantypes.QueryPacketCommitmentsResponse,
	error,
) {
	var res *chantypes.QueryPacketCommitmentsResponse
	err := i.client.CallContext(ctx, &res, queryPacketCommitmentsMethod, height, channelid, portid)
	if err != nil {
		return &chantypes.QueryPacketCommitmentsResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryPacketAcknowledgements(
	ctx context.Context,
	height uint32,
	channelid,
	portid string,
) (
	*chantypes.QueryPacketAcknowledgementsResponse,
	error,
) {
	var res *chantypes.QueryPacketAcknowledgementsResponse
	err := i.client.CallContext(ctx, &res, queryPacketAcknowledgementsMethod, height, channelid, portid)
	if err != nil {
		return &chantypes.QueryPacketAcknowledgementsResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryUnreceivedPackets(
	ctx context.Context,
	height uint32,
	channelid,
	portid string,
	seqs []uint64,
) (
	[]uint64, error,
) {
	var res []uint64
	err := i.client.CallContext(ctx, &res, queryUnreceivedPacketsMethod, height, channelid, portid, seqs)
	if err != nil {
		return []uint64{}, err
	}
	return res, nil
}

func (i IBC) QueryUnreceivedAcknowledgements(
	ctx context.Context,
	height uint32,
	channelid,
	portid string,
	seqs []uint64,
) (
	[]uint64,
	error,
) {
	var res []uint64
	err := i.client.CallContext(ctx, &res, queryUnreceivedAcknowledgementMethod, height, channelid, portid, seqs)
	if err != nil {
		return []uint64{}, err
	}
	return res, nil
}

func (i IBC) QueryNextSeqRecv(
	ctx context.Context,
	height uint32,
	channelid,
	portid string,
) (
	*chantypes.QueryNextSequenceReceiveResponse,
	error,
) {
	var res *chantypes.QueryNextSequenceReceiveResponse
	err := i.client.CallContext(ctx, &res, queryNextSeqRecvMethod, height, channelid, portid)
	if err != nil {
		return &chantypes.QueryNextSequenceReceiveResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryPacketCommitment(
	ctx context.Context,
	height int64,
	channelid,
	portid string,
) (
	*chantypes.QueryPacketCommitmentResponse,
	error,
) {
	var res *chantypes.QueryPacketCommitmentResponse
	err := i.client.CallContext(ctx, &res, queryPacketCommitmentMethod, height, channelid, portid)
	if err != nil {
		return &chantypes.QueryPacketCommitmentResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryPacketAcknowledgement(
	ctx context.Context,
	height uint32,
	channelid,
	portid string,
	seq uint64,
) (
	*chantypes.QueryPacketAcknowledgementResponse,
	error,
) {
	var res *chantypes.QueryPacketAcknowledgementResponse
	err := i.client.CallContext(ctx, &res, queryPacketAcknowledgementMethod, height, channelid, portid, seq)
	if err != nil {
		return &chantypes.QueryPacketAcknowledgementResponse{}, err
	}
	return res, nil
}

func (i IBC) QueryPacketReceipt(
	ctx context.Context,
	height uint32,
	channelid,
	portid string,
	seq uint64,
) (
	*chantypes.QueryPacketReceiptResponse,
	error,
) {
	var res *chantypes.QueryPacketReceiptResponse
	err := i.client.CallContext(ctx, &res, queryPacketReceiptMethod, height, channelid, portid, seq)
	if err != nil {
		return &chantypes.QueryPacketReceiptResponse{}, err
	}
	return res, nil
}
