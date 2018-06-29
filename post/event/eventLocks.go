package post

import "sync"

//GenericEventUpdateLock ...
var GenericEventUpdateLock sync.Mutex

//EventMemberLock ... PLEASE USE IF MODIFYING EVENT MEMBERS USING GENERIC UPDATE
var EventMemberLock sync.Mutex
