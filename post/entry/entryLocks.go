package post

import "sync"

var genericEntryUpdateLock sync.Mutex
var likeArrayLock sync.Mutex
var shareArrayLock sync.Mutex
var replyArrayLock sync.Mutex
