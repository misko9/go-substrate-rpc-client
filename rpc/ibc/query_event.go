package ibc

func (i IBC) QueryIbcEvents(
	blockNumbers []uint32,
) (
	interface{},
	error,
) {
	var res interface{}
	err := i.client.Call(&res, queryIbcEventsMethod, blockNumbers)
	if err != nil {
		return res, err
	}
	return res, nil
}
