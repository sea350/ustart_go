package uses

import (
	"log"

	getChat "github.com/sea350/ustart_go/get/chat"
	postChat "github.com/sea350/ustart_go/post/chat"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChatFirst ... Executes all necessary database interactions for a DM initiation
//returns the new convo id along with err
func ChatFirst(eclient *elastic.Client, msg types.Message, docID1 string, docID2 string) (string, error) {

	var eaves types.Eavesdropper
	eaves.Class = 1
	var eaves2 types.Eavesdropper
	eaves2.Class = 1
	var droppers = make(map[string]types.Eavesdropper)
	droppers[docID1] = eaves
	droppers[docID2] = eaves2
	var newConvo types.Conversation
	newConvo.Class = 1
	newConvo.Eavesdroppers = droppers

	convoID, err := postChat.IndexConvo(eclient, newConvo)
	if err != nil {

		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return ``, err
	}

	msg.ConversationID = convoID
	msgID, err := postChat.IndexMsg(eclient, msg)
	if err != nil {

		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return ``, err
	}

	err = postChat.UpdateConvo(eclient, convoID, "MessageArchive", []string{msgID})
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return convoID, err
	}

	pID, err := getChat.ProxyIDByUserID(eclient, docID1)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return convoID, err
	}
	err = postChat.AppendToProxy(eclient, pID, convoID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return convoID, err
	}

	pID, err = getChat.ProxyIDByUserID(eclient, docID2)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return convoID, err
	}
	err = postChat.AppendToProxy(eclient, pID, convoID)

	return convoID, err

}
