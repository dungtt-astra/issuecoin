package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "issuecoin/testutil/keeper"
	"issuecoin/x/issuecoin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IssuecoinKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
