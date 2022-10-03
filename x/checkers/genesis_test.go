package checkers_test

import (
	"testing"

	keepertest "github.com/hebx/checkers/testutil/keeper"
	"github.com/hebx/checkers/testutil/nullify"
	"github.com/hebx/checkers/x/checkers"
	"github.com/hebx/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: &types.SystemInfo{
		NextId: 20,
},
		SystemInfo: &types.SystemInfo{
		NextId: 47,
},
		SystemInfo: &types.SystemInfo{
		NextId: 53,
},
		SystemInfo: &types.SystemInfo{
		NextId: 96,
},
		SystemInfo: &types.SystemInfo{
		NextId: 72,
},
		SystemInfo: &types.SystemInfo{
		NextId: 76,
},
		SystemInfo: &types.SystemInfo{
		NextId: 34,
},
		SystemInfo: &types.SystemInfo{
		NextId: 58,
},
		SystemInfo: &types.SystemInfo{
		NextId: 10,
},
		SystemInfo: &types.SystemInfo{
		NextId: 82,
},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, genesisState)
	got := checkers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
// this line is used by starport scaffolding # genesis/test/assert
}
