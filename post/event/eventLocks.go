package post

import "sync"

//GenericEventUpdateLock ...
var GenericEventUpdateLock sync.Mutex

//EventMemberLock ... PLEASE USE IF MODIFYING EVENT MEMBERS USING GENERIC UPDATE
var EventMemberLock sync.Mutex

//EventGuestLock ... Use if modifying event guest using generic update
var EventGuestLock sync.Mutex

//EventGuestRequestLock ... Use if modifying event guest requests
var EventGuestRequestLock sync.Mutex

//EventFollowerLock ... a lock for modifying a event's follower array
var EventFollowerLock sync.Mutex
