package post

import "sync"

//ModifyMemberLock ... PLEASE USE IF MODIFYING MEMBERS USING GENERIC UPDATE
var ModifyMemberLock sync.Mutex

//GenericProjectUpdateLock ...
var GenericProjectUpdateLock sync.Mutex
