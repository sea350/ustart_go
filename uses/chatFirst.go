package uses

import (
	getChat "github.com/sea350/ustart_go/get/Chat"
	postChat "github.com/sea350/ustart_go/post/Chat"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChatFirst ... Executes all necessary database interactions for a DM initiation
func ChatFirst(eclient *elastic.Client, msg types.Message, docID1 string, docID2 string) error {

	eaves := types.Eavesdropper{Class: 1}

	droppers := make(map[string]types.Eavesdropper)
	droppers[docID1] = eaves
	droppers[docID2] = eaves
	newConvo := types.Conversation{Class: 1, Eavesdroppers: droppers}

	convoID, err := postChat.IndexConvo(eclient, newConvo)
	if err != nil {
		return err
	}

	msg.ConversationID = convoID
	msgID, err := postChat.IndexMsg(eclient, msg)
	if err != nil {
		return err
	}

	err = postChat.UpdateConvo(eclient, convoID, "MessageArchive", []string{msgID})
	if err != nil {
		return err
	}

	pID, err := getChat.ProxyIDByUserID(eclient, docID1)
	if err != nil {
		return err
	}
	err = postChat.AppendToProxy(eclient, pID, convoID)
	if err != nil {
		return err
	}

	pID, err := getChat.ProxyIDByUserID(eclient, docID2)
	if err != nil {
		return err
	}
	err = postChat.AppendToProxy(eclient, pID, convoID)

	return err

}
