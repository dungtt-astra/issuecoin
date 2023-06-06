package keeper

import (
	"context"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"issuecoin/x/issuecoin/types"
)

func (k msgServer) Mintcoin(goCtx context.Context, msg *types.MsgMintcoin) (*types.MsgMintcoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	log.Println("Start Mintcoin..")
	_ = ctx
	err := k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{*msg.Coin})
	if err != nil {
		log.Println("Start Mintcoin... err:", err)
		return nil, err
	}

	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, sdk.Coins{*msg.Coin})
	if err != nil {
		return nil, err
	}
	return &types.MsgMintcoinResponse{}, nil

}
