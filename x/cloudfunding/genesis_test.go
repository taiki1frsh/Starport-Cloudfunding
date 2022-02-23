package cloudfunding_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/taiki-furu/cloudfunding/testutil/keeper"
	"github.com/taiki-furu/cloudfunding/testutil/nullify"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ProjectList: []types.Project{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ProjectCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CloudfundingKeeper(t)
	cloudfunding.InitGenesis(ctx, *k, genesisState)
	got := cloudfunding.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProjectList, got.ProjectList)
	require.Equal(t, genesisState.ProjectCount, got.ProjectCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
