package ibc

import "github.com/ComposableFi/go-substrate-rpc-client/v4/types"

func (i IBC) QueryIbcEvents(
	blockNumbers []types.BlockNumberOrHash,
) (
	map[string][]interface{},
	error,
) {
	var res map[string][]interface{}
	err := i.client.Call(&res, queryIbcEventsMethod, blockNumbers)
	if err != nil {
		return res, err
	}
	return res, nil
}
