package elysium_test

import (
	"testing"

	keepertest "elysium/testutil/keeper"
	"elysium/testutil/nullify"
	"elysium/x/elysium"
	"elysium/x/elysium/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ElysiumKeeper(t)
	elysium.InitGenesis(ctx, *k, genesisState)
	got := elysium.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
