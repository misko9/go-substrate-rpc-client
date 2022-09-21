package ibc

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (i IBC) QueryBalanceWithAddress(
	ctx context.Context,
	addr []byte,
) (
	sdk.Coins,
	error,
) {
	var res sdk.Coins
	err := i.client.CallContext(ctx, &res, queryBalanceWithAddressMethod, addr)
	if err != nil {
		return sdk.Coins{}, err
	}
	return res, nil
}
