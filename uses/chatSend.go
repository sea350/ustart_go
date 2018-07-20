package uses

import (
	"errors"
	"log"

	getChat "github.com/sea350/ustart_go/get/chat"
	postChat "github.com/sea350/ustart_go/post/chat"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChatSend ... Executes all necessary database interactions for a message to a chat
func ChatSend(eclient *elastic.Client, msg types.Message) ([]string, error) {

	var notifyThese []string

	convo, err := getChat.ConvoByID(eclient, msg.ConversationID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return notifyThese, err
	}

	_, exists := convo.Eavesdroppers[msg.SenderID]
	if !exists {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(errors.New("THIS USER IS NOT PART OF THE CONVERSATION"))
		return notifyThese, errors.New("THIS USER IS NOT PART OF THE CONVERSATION")
	}

	msgID, err := postChat.IndexMsg(eclient, msg)
	if err != nil {
		return notifyThese, err
	}

	err = postChat.AppendMessageIDToConversation(eclient, msg.ConversationID, msgID)
	for eaverID := range convo.Eavesdroppers {
		pID, err := getChat.ProxyIDByUserID(eclient, eaverID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return notifyThese, err
		}
		err = postChat.AppendToProxy(eclient, pID, msg.ConversationID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return notifyThese, err
		}

		notifyThese = append(notifyThese, eaverID)

	}

	return notifyThese, err

}
