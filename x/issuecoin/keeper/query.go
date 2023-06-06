package keeper

import (
	"issuecoin/x/issuecoin/types"
)

var _ types.QueryServer = Keeper{}
