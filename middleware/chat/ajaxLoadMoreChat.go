package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/sea350/ustart_go/uses"

	"github.com/sea350/ustart_go/middleware/client"
)

//AjaxLoadMoreChat ... loads more chat
//Designed for ajax
func AjaxLoadMoreChat(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	chatID := r.FormValue("chatID")
	idxStr := r.FormValue("index")
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	valid, actualChatID, _, err := uses.ChatVerifyURL(client.Eclient, chatID, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	if !valid {
		return
	}

	newIdx, msgs, err := uses.ChatLoad(client.Eclient, actualChatID, 0, idx)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	sendThis := make(map[string]interface{})
	sendThis["Index"] = newIdx
	sendThis["Messages"] = msgs

	data, err := json.Marshal(sendThis)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	fmt.Fprintln(w, string(data))
}
