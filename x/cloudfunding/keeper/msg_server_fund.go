package keeper

import (
	"context"

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

	amt, err := sdk.ParseCoinsNormalized(msg.Amt)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid coin")
	}
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, funder, types.ModuleName, amt)

	k.SetProject(ctx, project)
	return &types.MsgFundResponse{}, nil
}
