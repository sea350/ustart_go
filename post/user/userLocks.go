package post

import "sync"

//procLock ...
var procLock sync.Mutex

//likeLock ...
var likeLock sync.Mutex

//colleagueLock ...
var colleagueLock sync.Mutex

//followLock ...
var followLock sync.Mutex

//projectLock ...
var projectLock sync.Mutex

//blockLock ...
var blockLock sync.Mutex

//tagLock ...
var tagLock sync.Mutex

//entryLock ...
var entryLock sync.Mutex
