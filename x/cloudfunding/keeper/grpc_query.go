package keeper

import (
	"github.com/taiki-furu/cloudfunding/x/cloudfunding/types"
)

var _ types.QueryServer = Keeper{}
