package post

import "sync"

//GenericEventUpdateLock ...
var GenericEventUpdateLock sync.Mutex

//EventMemberLock ... PLEASE USE IF MODIFYING EVENT MEMBERS USING GENERIC UPDATE
var EventMemberLock sync.Mutex

//EventGuestLock ... Use if modifying event guest using generic update
var EventGuestLock sync.Mutex
