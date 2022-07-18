package ibc

import "github.com/ComposableFi/go-substrate-rpc-client/v4/client"

const (
	generateConnectionHandshakeProof     = "ibc_generateConnectionHandshakeProof"
	queryAcknowledgements                = "ibc_queryAcknowledgements"
	queryBalanceWithAddressMethod        = "ibc_queryBalanceWithAddress"
	queryChannelClientMethod             = "ibc_queryChannelClient"
	queryChannelsMethod                  = "ibc_queryChannels"
	queryClientStateMethod               = "ibc_queryClientState"
	queryClientConsensusStateMethod      = "ibc_queryClientConsensusState"
	queryClientsMethod                   = "ibc_queryClients"
	queryConnectionMethod                = "ibc_queryConnection"
	queryConnectionsMethod               = "ibc_queryConnections"
	queryConnectionChannelsMethod        = "ibc_queryConnectionChannels"
	queryConnectionUsingClientMethod     = "ibc_queryConnectionUsingClient"
	queryConsensusStateMethod            = "ibc_queryConsensusState"
	queryDenomTraceMethod                = "ibc_queryDenomTrace"
	queryDenomTracesMethod               = "ibc_queryDenomTraces"
	queryLatestHeight                    = "ibc_queryLatestHeight"
	queryNextSeqRecvMethod               = "ibc_queryNextSeqRecv"
	queryNewlyCreatedClients             = "ibc_queryNewlyCreatedClients"
	queryPacketsMethod                   = "ibc_queryPackets"
	queryPacketCommitmentsMethod         = "ibc_queryPacketCommitments"
	queryPacketAcknowledgementsMethod    = "ibc_queryPacketAcknowledgements"
	queryPacketCommitmentMethod          = "ibc_queryPacketCommitment"
	queryPacketAcknowledgementMethod     = "ibc_queryPacketAcknowledgement"
	queryPacketReceiptMethod             = "ibc_queryPacketReceipt"
	queryProof                           = "ibc_queryProof"
	queryUnreceivedAcknowledgementMethod = "ibc_queryUnreceivedAcknowledgement"
	queryUnreceivedPacketsMethod         = "ibc_queryUnreceivedPackets"
	queryUpgradedClientMethod            = "ibc_queryUpgradedClient"
	queryUpgradedConnectionStateMethod   = "ibc_queryUpgradedConnectionState"
)

// IBC exposes methods for retrieval of chain data
type IBC struct {
	client client.Client
}

// NewIBC creates a new IBC struct
func NewIBC(cl client.Client) *IBC {
	return &IBC{cl}
}
