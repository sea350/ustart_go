package uses

import (
	"log"

	getChat "github.com/sea350/ustart_go/get/chat"
	postChat "github.com/sea350/ustart_go/post/chat"
	"github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ChatFirst ... Executes all necessary database interactions for a DM initiation
//returns the new convo id along with err
func ChatFirst(eclient *elastic.Client, msg types.Message, docID1 string, docID2 string) (string, error) {

	var newConvo = types.Conversation{Class: 1}

	droppers := []types.Eavesdropper{}
	droppers = append(droppers, types.Eavesdropper{Class: 1, DocID: docID1})
	if docID1 != docID2 {
		droppers = append(droppers, types.Eavesdropper{Class: 1, DocID: docID2})
		newConvo.Size = 2
	} else {
		newConvo.Size = 1
	}

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
	err = postChat.AppendToProxy(eclient, pID, convoID, false)
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
	err = postChat.AppendToProxy(eclient, pID, convoID, false)

	return convoID, err

}
