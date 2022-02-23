package cloudfunding

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/keeper"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the project
	for _, elem := range genState.ProjectList {
		k.SetProject(ctx, elem)
	}

	// Set project count
	k.SetProjectCount(ctx, genState.ProjectCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ProjectList = k.GetAllProject(ctx)
	genesis.ProjectCount = k.GetProjectCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
