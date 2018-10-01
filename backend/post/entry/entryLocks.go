package post

import "sync"

//GenericEntryUpdateLock ...
var GenericEntryUpdateLock sync.Mutex

//LikeArrayLock ...
var LikeArrayLock sync.Mutex

//ShareArrayLock ...
var ShareArrayLock sync.Mutex

//ReplyArrayLock ...
var ReplyArrayLock sync.Mutex

//EntryLock ...
var EntryLock sync.Mutex
