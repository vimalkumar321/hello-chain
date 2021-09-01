package keeper

import (
	"github.com/example/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}
