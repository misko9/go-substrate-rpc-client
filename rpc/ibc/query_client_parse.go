package ibc

import (
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

type IdentifiedClientStates []IdentifiedClientState

type IdentifiedClientState struct {
	// client identifier
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" yaml:"client_id"`
	// client state
	ClientState *Any `protobuf:"bytes,2,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty" yaml:"client_state"`
}

func parseIdentifiedClientStates(ics IdentifiedClientStates) clienttypes.IdentifiedClientStates {
	var clientStates clienttypes.IdentifiedClientStates
	for i := 0; i < len(ics); i++ {
		cs := parseAny(ics[i].ClientState)
		clientStates = append(clientStates, clienttypes.IdentifiedClientState{
			ClientState: cs,
			ClientId:    ics[i].ClientId,
		})
	}
	return clientStates
}

type QueryClientStateResponse struct {
	// client state associated with the request identifier
	ClientState *Any `protobuf:"bytes,1,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty"`
	// merkle proof of existence
	Proof []byte `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
	// height at which the proof was retrieved
	ProofHeight clienttypes.Height `protobuf:"bytes,3,opt,name=proof_height,json=proofHeight,proto3" json:"proof_height"`
}

func parseQueryClientStateResponse(csr QueryClientStateResponse) clienttypes.QueryClientStateResponse {
	return clienttypes.QueryClientStateResponse{
		ClientState: parseAny(csr.ClientState),
		Proof: csr.Proof,
		ProofHeight: csr.ProofHeight,
	}
}
