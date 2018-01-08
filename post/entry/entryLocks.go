package post

import "sync"

var GenericEntryUpdateLock sync.Mutex
var LikeArrayLock sync.Mutex
var ShareArrayLock sync.Mutex
var ReplyArrayLock sync.Mutex
var EntryLock sync.Mutex
