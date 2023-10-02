package keeper

import (
	"cool-chain/x/coolchain/types"
)

var _ types.QueryServer = Keeper{}
