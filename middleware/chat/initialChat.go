package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	getChat "github.com/sea350/ustart_go/get/chat"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"

	"github.com/sea350/ustart_go/middleware/client"
)

//InitialChat ... crawling in the 90s
//Designed for ajax
func InitialChat(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	chatURL := r.FormValue("chatUrl")

	valid, actualChatID, otherUsr, err := uses.ChatVerifyURL(client.Eclient, chatURL, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	if !valid {
		return
	}

	heads := make(map[string]types.FloatingHead)

	if actualChatID != `` {
		_, msgs, err := uses.ChatLoad(client.Eclient, actualChatID, 0, 30)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			return
		}

		// data, err := json.Marshal(size)
		// if err != nil {
		// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
		// 	dir, _ := os.Getwd()
		// 	log.Println(dir, err)
		// }
		// fmt.Fprintln(w, string(data))

		data, err := json.Marshal(msgs)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		fmt.Fprintln(w, string(data))

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
			heads[chat.Eavesdroppers[idx].DocID] = head
		}

	} else {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, otherUsr)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		heads[otherUsr] = head

		head, err = uses.ConvertUserToFloatingHead(client.Eclient, docID.(string))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		heads[docID.(string)] = head
	}

	// data, err := json.Marshal(heads)
	// if err != nil {
	// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 	dir, _ := os.Getwd()
	// 	log.Println(dir, err)
	// }
	// fmt.Fprintln(w, string(data))
}
