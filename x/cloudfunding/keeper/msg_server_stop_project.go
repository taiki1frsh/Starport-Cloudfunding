package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

func (k msgServer) StopProject(goCtx context.Context, msg *types.MsgStopProject) (*types.MsgStopProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	project, found := k.GetProject(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
	}

	if msg.Creator != project.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "you are not creator of this project")
	}

	blockheight := ctx.BlockHeight()
	deadline, _ := strconv.ParseInt(project.Deadline, 10, 64)
	if blockheight < deadline {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "still in funding period")
	}

	collected := project.Collected + "token"
	collectedCoins, _ := sdk.ParseCoinsNormalized(collected)

	// In any case, collected amount goes to creator address
	creator, _ := sdk.AccAddressFromBech32(project.Creator)
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, collectedCoins)
	if sdkError != nil {
		return nil, sdkError
	}

	project.State = "ended"

	return &types.MsgStopProjectResponse{}, nil
}
