package cache

import (
	"sync"

	tmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const (
	DefaultCacheSize = 1000000
)

var (
	BlockCache      []*tmrpctypes.ResultBlockResults
	BlockCacheMutex = sync.Mutex{}
)

func init() {
	blockCache := make([]*tmrpctypes.ResultBlockResults, DefaultCacheSize)
	for i := 0; i < DefaultCacheSize; i++ {
		blockCache[i] = nil
	}

	BlockCache = blockCache
}
