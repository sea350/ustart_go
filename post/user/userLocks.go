package post

import "sync"

//procLock ...
var ProcLock sync.Mutex

//likeLock ...
var LikeLock sync.Mutex

//colleagueLock ...
var ColleagueLock sync.Mutex

//followLock ...
var FollowLock sync.Mutex

//projectLock ...
var ProjectLock sync.Mutex

//blockLock ...
var BlockLock sync.Mutex

//tagLock ...
var TagLock sync.Mutex

//entryLock ...
var EntryLock sync.Mutex
