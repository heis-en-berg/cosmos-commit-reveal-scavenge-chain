package keeper

import (
	"github.com/heis-en-berg/scavenge/x/scavenge/types"
)

var _ types.QueryServer = Keeper{}
