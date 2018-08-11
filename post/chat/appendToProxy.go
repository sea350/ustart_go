package post

import (
	"context"
	"errors"
	"log"

	get "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendToProxy ... appends a new conversation state OR brings a certain conversation state to the back of the list
//needs its own lock for concurrency control
func AppendToProxy(eclient *elastic.Client, proxyID string, conversationID string, seen bool) error {

	ctx := context.Background()

	AppendToProxyLock.Lock()
	defer AppendToProxyLock.Unlock()

	exists, err := eclient.IndexExists(globals.ProxyMsgIndex).Do(ctx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}
	if !exists {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(errors.New("Index does not exist"))
		return errors.New("Index does not exist")
	}

	proxy, err := get.ProxyMsgByID(eclient, proxyID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return err
	}

	var temp types.ConversationState
	for i := len(proxy.Conversations) - 1; i >= 0; i-- {
		if proxy.Conversations[i].ConvoID == conversationID {
			temp = proxy.Conversations[i]
			proxy.Conversations = append(proxy.Conversations[:i], proxy.Conversations[i+1:]...)
			break
		}
	}

	if temp.ConvoID == `` { //adding a new convo
		temp.ConvoID = conversationID
		proxy.NumUnread++
	} else if !temp.Read && seen {
		temp.Read = false
		proxy.NumUnread--
	}
	if temp.Read && !seen {
		temp.Read = false
		proxy.NumUnread++
	}
	proxy.Conversations = append(proxy.Conversations, temp)

	err = ReindexProxyMsg(eclient, proxyID, proxy)

	return err
}
