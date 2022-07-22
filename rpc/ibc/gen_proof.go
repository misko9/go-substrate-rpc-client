package ibc

import "github.com/ComposableFi/go-substrate-rpc-client/v4/types"

func (i IBC) GenerateConnectionHandshakeProof(
	height uint32,
	clientId string,
	connId string,
) (
	types.ConnHandshakeProof,
	error,
) {
	var res types.ConnHandshakeProof
	err := i.client.Call(&res, generateConnectionHandshakeProofMethod)
	if err != nil {
		return types.ConnHandshakeProof{}, err
	}
	return res, nil
}
