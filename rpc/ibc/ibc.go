package ibc

import "github.com/ComposableFi/go-substrate-rpc-client/v4/client"

const (
	generateConnectionHandshakeProofMethod = "ibc_generateConnectionHandshakeProof"
	queryAcknowledgementsMethod            = "ibc_queryAcknowledgements"
	queryBalanceWithAddressMethod          = "ibc_queryBalanceWithAddress"
	queryChannelClientMethod               = "ibc_queryChannelClient"
	queryChannelsMethod                    = "ibc_queryChannels"
	queryClientStateMethod                 = "ibc_queryClientState"
	queryClientConsensusStateMethod        = "ibc_queryClientConsensusState"
	queryClientsMethod                     = "ibc_queryClients"
	queryConnectionMethod                  = "ibc_queryConnection"
	queryConnectionsMethod                 = "ibc_queryConnections"
	queryConnectionChannelsMethod          = "ibc_queryConnectionChannels"
	queryConnectionUsingClientMethod       = "ibc_queryConnectionUsingClient"
	queryConsensusStateMethod              = "ibc_queryConsensusState"
	queryDenomTraceMethod                  = "ibc_queryDenomTrace"
	queryDenomTracesMethod                 = "ibc_queryDenomTraces"
	queryIbcEventsMethod                   = "ibc_queryIbcEvents"
	queryLatestHeightMethod                = "ibc_queryLatestHeight"
	queryNextSeqRecvMethod                 = "ibc_queryNextSeqRecv"
	queryNewlyCreatedClientsMethod         = "ibc_queryNewlyCreatedClients"
	queryPacketsMethod                     = "ibc_queryPackets"
	queryPacketCommitmentsMethod           = "ibc_queryPacketCommitments"
	queryPacketAcknowledgementsMethod      = "ibc_queryPacketAcknowledgements"
	queryPacketCommitmentMethod            = "ibc_queryPacketCommitment"
	queryPacketAcknowledgementMethod       = "ibc_queryPacketAcknowledgement"
	queryPacketReceiptMethod               = "ibc_queryPacketReceipt"
	queryProofMethod                       = "ibc_queryProof"
	queryUnreceivedAcknowledgementMethod   = "ibc_queryUnreceivedAcknowledgement"
	queryUnreceivedPacketsMethod           = "ibc_queryUnreceivedPackets"
	queryUpgradedClientMethod              = "ibc_queryUpgradedClient"
	queryUpgradedConnectionStateMethod     = "ibc_queryUpgradedConnectionState"
)

// IBC exposes methods for retrieval of chain data
type IBC struct {
	client client.Client
}

// NewIBC creates a new IBC struct
func NewIBC(cl client.Client) *IBC {
	return &IBC{cl}
}
