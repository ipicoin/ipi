package ipi_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ipi/testutil/keeper"
	"ipi/testutil/nullify"
	"ipi/x/ipi"
	"ipi/x/ipi/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IpiKeeper(t)
	ipi.InitGenesis(ctx, *k, genesisState)
	got := ipi.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
