package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var project = types.Project{
		Target:      msg.Target,
		Collected:   "0",
		Deadline:    "nil",
		Description: msg.Description,
		State:       "undergoing",
		Creator:     msg.Creator,
	}

	blockHeight := ctx.BlockHeight()
	deadlineInt, err := strconv.ParseInt(project.Deadline, 10, 64)
	if err != nil {
		panic(err)
	}
	rightDeadline := blockHeight + deadlineInt
	project.Deadline = strconv.Itoa(int(rightDeadline))

	k.SetProject(ctx, project)

	return &types.MsgCreateProjectResponse{}, nil
}
