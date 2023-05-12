package keeper_test

import (
	"context"
	"testing"

	keepertest "elysium/testutil/keeper"
	"elysium/x/elysium/keeper"
	"elysium/x/elysium/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ElysiumKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
