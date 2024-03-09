package cache

import (
	"sync"

	tmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const (
	DefaultCacheSize = 1000000
)

var (
	BlockCache      = make(map[int64]*tmrpctypes.ResultBlockResults, DefaultCacheSize)
	BlockCacheMutex = sync.Mutex{}
)
