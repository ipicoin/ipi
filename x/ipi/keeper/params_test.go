package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ipi/testutil/keeper"
	"ipi/x/ipi/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IpiKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
