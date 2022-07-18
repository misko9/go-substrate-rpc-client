package ibc

func (i IBC) QueryLatestHeight() (
	uint64,
	error,
) {
	var res uint64
	err := i.client.Call(&res, queryLatestHeight)
	if err != nil {
		return 0, err
	}
	return res, nil
}
