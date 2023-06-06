package issuecoin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "issuecoin/testutil/keeper"
	"issuecoin/testutil/nullify"
	"issuecoin/x/issuecoin"
	"issuecoin/x/issuecoin/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IssuecoinKeeper(t)
	issuecoin.InitGenesis(ctx, *k, genesisState)
	got := issuecoin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
