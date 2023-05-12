package keeper_test

import (
	"testing"

	testkeeper "elysium/testutil/keeper"
	"elysium/x/elysium/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ElysiumKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
