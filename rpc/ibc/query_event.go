package ibc

import "github.com/ComposableFi/go-substrate-rpc-client/v4/types"

func (i IBC) QueryIbcEvents(
	blockNumbers []types.BlockNumberOrHash,
) (
	types.IBCEventsQueryResult,
	error,
) {
	var res types.IBCEventsQueryResult
	err := i.client.Call(&res, queryIbcEventsMethod, blockNumbers)
	if err != nil {
		return res, err
	}
	return res, nil
}
