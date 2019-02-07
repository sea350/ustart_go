package uses

import (
	"log"

	getChat "github.com/sea350/ustart_go/get/chat"
	getUser "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	postUser "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChatAggregateNotifications ... Executes all necessary database interactions to pull chat notifs
func ChatAggregateNotifications(eclient *elastic.Client, usrID string, index int) ([]types.FloatingHead, int, error) {

	var notifs []types.FloatingHead
	var numUnread int

	usr, err := getUser.UserByID(client.Eclient, usrID)
	if err != nil {
		return notifs, numUnread, err
	}

	if usr.ProxyMessagesID == `` {
		prox, _ := getChat.ProxyIDByUserID(eclient, usrID)
		if prox != `` { //resync
			err = postUser.UpdateUser(client.Eclient, usrID, "ProxyMessagesID", prox)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				return notifs, numUnread, err
			}
		} else {
			newProxy := types.ProxyMessages{DocID: usrID, Class: 1}
			proxyID, err := postChat.IndexProxyMsg(client.Eclient, newProxy)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				return notifs, numUnread, err
			}
			err = postUser.UpdateUser(client.Eclient, usrID, "ProxyMessagesID", proxyID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				return notifs, numUnread, err
			}
		}
	} else {
		proxy, err := getChat.ProxyMsgByID(client.Eclient, usr.ProxyMessagesID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return notifs, numUnread, err
		}

		numUnread = proxy.NumUnread

		length := len(proxy.Conversations)
		startAt := (length - 1) - index
		if startAt < 0 {
			return notifs, numUnread, nil
		}
		endAt := (length - 1) - (index + 20)
		if endAt < 0 {
			endAt = 0
		}

		for i := startAt; i >= endAt; i-- {
			head, err := ConvertChatToFloatingHead(client.Eclient, proxy.Conversations[i].ConvoID, usrID)
			if err == nil {
				head.Read = proxy.Conversations[i].Read
				notifs = append(notifs, head)
			}
		}
	}

	return notifs, numUnread, err
}
