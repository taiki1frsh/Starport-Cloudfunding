package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/taiki-furu/cloudfunding/testutil/keeper"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/keeper"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CloudfundingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
