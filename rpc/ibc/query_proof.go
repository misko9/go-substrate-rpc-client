package ibc

import "github.com/ComposableFi/go-substrate-rpc-client/v4/types"

func (i IBC) QueryProof(
	height uint32,
	keys [][]byte) (
	types.Proof,
	error,
) {
	var res types.Proof
	err := i.client.Call(&res, queryProofMethod, height, keys)
	if err != nil {
		return types.Proof{}, err
	}
	return res, nil
}
