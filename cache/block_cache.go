package cache

import "sync"

var BlockCache = sync.Map{}
var Mutex = sync.Mutex{}
