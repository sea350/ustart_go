package post

import "sync"

//ProcLock ...
var ProcLock sync.Mutex

//LikeLock ...
var LikeLock sync.Mutex

//ColleagueLock ...
var ColleagueLock sync.Mutex

//FollowLock ...
var FollowLock sync.Mutex

//ProjectLock ...
var ProjectLock sync.Mutex

//BlockLock ...
var BlockLock sync.Mutex

//TagLock ...
var TagLock sync.Mutex

//EntryLock ...
var EntryLock sync.Mutex
