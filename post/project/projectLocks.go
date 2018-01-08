package post

import "sync"

//ModifyMemberLock ... PLEASE USE IF MODIFYING MEMBERS USING GENERIC UPDATE
var ModifyMemberLock sync.Mutex

//genericProjectUpdateLock ...
var GenericProjectUpdateLock sync.Mutex
