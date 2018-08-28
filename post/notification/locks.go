package post

import "sync"

//AppendMessageLock ... whenever you need to add a message remember to lock to avoid concurrency overwrites
var AppendMessageLock sync.Mutex

//AppendToProxyLock ... multiple active chats means frequently sending an active converstationState to the back of the map
var AppendToProxyLock sync.Mutex

//ModifyUnseen ... if you are changing unseen you need to lock
var ModifyUnseen sync.Mutex

//ModifyUnread ... if you are changing unread you need to lock
var ModifyUnread sync.Mutex
