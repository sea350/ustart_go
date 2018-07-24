package uses

import (
	"errors"
	"log"

	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChatVerifyURL ... Executes all necessary database interactions to verify the existance of and user's acccess to a conversation
//returns if the chat is valid, the actual id of conversation, docID of the second dmer if dm, and error
func ChatVerifyURL(eclient *elastic.Client, url string, viewerID string) (bool, string, string, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("debug text: " + url)

	if len(url) > 0 {
		if url[:1] == "@" {
			//This means its a DM
			targetID, err := get.IDByUsername(client.Eclient, url[1:])
			if err != nil {
				//invalid username
				return false, ``, ``, err
			}
			//checks if the dm already exists
			exists, id, err := getChat.DMExists(client.Eclient, targetID, viewerID)
			if err != nil {
				//there was a problem finding the dm
				return exists, id, ``, err
			}
			if exists {
				return exists, id, targetID, err
			}
			//this means a new DM is being opened, technically still a valid chat but no existing archive yet
			return true, ``, targetID, err

		}
		//the url is the actual chat ID
		convo, err := getChat.ConvoByID(client.Eclient, url)
		if err != nil {
			return false, ``, ``, err
		}

		var exists bool
		for i := range convo.Eavesdroppers {
			if convo.Eavesdroppers[i].DocID == viewerID {
				exists = true
				break
			}
		}
		if !exists {
			return false, ``, ``, errors.New("THIS USER IS NOT PART OF THE CONVERSATION")
		}

	}

	return true, url, ``, nil
}
