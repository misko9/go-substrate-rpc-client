package ibc

import (
	"github.com/ComposableFi/go-substrate-rpc-client/v4/client"
	sdktypes "github.com/cosmos/cosmos-sdk/codec/types"
)

const (
	queryBalanceWithAddressMethod        = "ibc_queryBalanceWithAddress"
	queryChannelClientMethod             = "ibc_queryChannelClient"
	queryConnectionChannelsMethod        = "ibc_queryConnectionChannels"
	queryChannelsMethod                  = "ibc_queryChannels"
	queryClientStateMethod               = "ibc_queryClientState"
	queryClientConsensusStateMethod      = "ibc_queryClientConsensusState"
	queryUpgradedClientMethod            = "ibc_queryUpgradedClient"
	queryUpgradedConnectionStateMethod   = "ibc_queryUpgradedConnectionState"
	queryClientsMethod                   = "ibc_queryClients"
	queryConnectionMethod                = "ibc_queryConnection"
	queryConnectionsMethod               = "ibc_queryConnections"
	queryConnectionUsingClientMethod     = "ibc_queryConnectionUsingClient"
	queryConsensusStateMethod            = "ibc_queryConsensusState"
	queryDenomTraceMethod                = "ibc_queryDenomTrace"
	queryDenomTracesMethod               = "ibc_queryDenomTraces"
	queryPacketsMethod                   = "ibc_queryPackets"
	queryPacketCommitmentsMethod         = "ibc_queryPacketCommitments"
	queryPacketAcknowledgementsMethod    = "ibc_queryPacketAcknowledgements"
	queryUnreceivedPacketsMethod         = "ibc_queryUnreceivedPackets"
	queryUnreceivedAcknowledgementMethod = "ibc_queryUnreceivedAcknowledgement"
	queryNextSeqRecvMethod               = "ibc_queryNextSeqRecv"
	queryPacketCommitmentMethod          = "ibc_queryPacketCommitment"
	queryPacketAcknowledgementMethod     = "ibc_queryPacketAcknowledgement"
	queryPacketReceiptMethod             = "ibc_queryPacketReceipt"
	queryNewlyCreatedClientsMethod       = "ibc_queryNewlyCreatedClients"
)

// IBC exposes methods for retrieval of chain data
type IBC struct {
	client client.Client
}

// NewIBC creates a new IBC struct
func NewIBC(cl client.Client) *IBC {
	return &IBC{cl}
}

// TODO: these custom Any and Client types need to be kept somewhere cleaner
type Any struct {
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Must be a valid serialized protocol buffer of the above specified type.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func parseAny(any *Any) *sdktypes.Any {
	return &sdktypes.Any{
		TypeUrl: any.TypeUrl,
		Value: any.Value,
	}
}
