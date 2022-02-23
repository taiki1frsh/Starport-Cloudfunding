package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

func (k msgServer) Fund(goCtx context.Context, msg *types.MsgFund) (*types.MsgFundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	project, found := k.GetProject(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
	}

	funder, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "incorrect address")
	}

	// Check if funding is available
	blockheight := ctx.BlockHeight()
	deadline, _ := strconv.ParseInt(project.Deadline, 10, 64)
	if blockheight > deadline {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "funding period is over")
	}

	tokenSuffix := "token"
	amt := msg.Amt + tokenSuffix
	amount, err := sdk.ParseCoinsNormalized(amt)
	if err != nil {
		panic(err)
	}

	collectedInt, _ := strconv.ParseInt(project.Collected, 10, 64)
	amtInt, _ := strconv.ParseInt(msg.Amt, 10, 64)
	collectedInt += amtInt
	project.Collected = strconv.Itoa(int(collectedInt))

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, funder, types.ModuleName, amount)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetProject(ctx, project)
	return &types.MsgFundResponse{}, nil
}
