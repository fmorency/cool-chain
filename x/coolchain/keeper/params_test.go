package keeper_test

import (
	"testing"

	testkeeper "cool-chain/testutil/keeper"
	"cool-chain/x/coolchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CoolchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
