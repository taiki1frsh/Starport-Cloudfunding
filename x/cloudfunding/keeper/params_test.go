package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/taiki-furu/cloudfunding/testutil/keeper"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CloudfundingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
