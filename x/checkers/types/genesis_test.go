package types_test

import (
	"testing"

	"github.com/hebx/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{

				SystemInfo: &types.SystemInfo{
		NextId: 4,
},
SystemInfo: &types.SystemInfo{
		NextId: 27,
},
SystemInfo: &types.SystemInfo{
		NextId: 19,
},
SystemInfo: &types.SystemInfo{
		NextId: 30,
},
SystemInfo: &types.SystemInfo{
		NextId: 3,
},
SystemInfo: &types.SystemInfo{
		NextId: 12,
},
SystemInfo: &types.SystemInfo{
		NextId: 31,
},
SystemInfo: &types.SystemInfo{
		NextId: 83,
},
SystemInfo: &types.SystemInfo{
		NextId: 43,
},
SystemInfo: &types.SystemInfo{
		NextId: 42,
},
// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
