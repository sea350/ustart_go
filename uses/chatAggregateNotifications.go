package uses

import (
	"log"
	"os"

	getChat "github.com/sea350/ustart_go/get/chat"
	getUser "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	postUser "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChatAggregateNotifications ... Executes all necessary database interactions to pull chat notifs
func ChatAggregateNotifications(eclient *elastic.Client, usrID string) ([]types.FloatingHead, error) {

	var notifs []types.FloatingHead

	usr, err := getUser.UserByID(client.Eclient, usrID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		return notifs, err
	}

	if usr.ProxyMessagesID == `` {
		newProxy := types.ProxyMessages{DocID: usrID, Class: 1}
		proxyID, err := postChat.IndexProxyMsg(client.Eclient, newProxy)
		err = postUser.UpdateUser(client.Eclient, usrID, "ProxyMesssagesID", proxyID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return notifs, err
		}
	} else {
		proxy, err := getChat.ProxyMsgByID(client.Eclient, usr.ProxyMessagesID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return notifs, err
		}

		for key := range proxy.Conversations {
			head, err := ConvertChatToFloatingHead(client.Eclient, key, usrID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
			}
			notifs = append(notifs, head)
		}
	}

	return notifs, err
}
