package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"

	getChat "github.com/sea350/ustart_go/get/chat"
	"github.com/sea350/ustart_go/middleware/client"
)

type chatAggregate struct {
	Index         int                           `json:"Index"`
	Messages      []types.Message               `json:"Messages"`
	Eavesdroppers map[string]types.FloatingHead `json:"Eavesdroppers"`
}

//InitialChat ... crawling in the 90s
//Designed for ajax
func InitialChat(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	chatURL := r.FormValue("chatUrl")

	var agg chatAggregate

	valid, actualChatID, otherUsr, err := uses.ChatVerifyURL(client.Eclient, chatURL, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	if !valid {
		return
	}

	if actualChatID != `` {
		_, agg.Messages, err = uses.ChatLoad(client.Eclient, actualChatID, 0, 50)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}

		chat, err := getChat.ConvoByID(client.Eclient, actualChatID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return
		}

		for idx := range chat.Eavesdroppers {
			head, err := uses.ConvertUserToFloatingHead(client.Eclient, chat.Eavesdroppers[idx].DocID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
			}
			agg.Eavesdroppers[chat.Eavesdroppers[idx].DocID] = head
		}

	} else {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, otherUsr)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		agg.Eavesdroppers[otherUsr] = head

		head, err = uses.ConvertUserToFloatingHead(client.Eclient, docID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		agg.Eavesdroppers[docID.(string)] = head
	}

	data, err := json.Marshal(agg)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	fmt.Fprintln(w, string(data))
}
