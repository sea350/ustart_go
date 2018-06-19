package post

import "sync"

//ModifyMemberLock ... PLEASE USE IF MODIFYING MEMBERS USING GENERIC UPDATE
var ModifyMemberLock sync.Mutex

//GenericProjectUpdateLock ...
var GenericProjectUpdateLock sync.Mutex

//FollowerLock ... a lock for modifying a project's follower array
var FollowerLock sync.Mutex
