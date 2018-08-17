package post

import "sync"

//ProcLock ...
var FollowerLock sync.Mutex

//FollowingLock
var FollowingLock sync.Mutex
