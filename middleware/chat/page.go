package chat

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/sea350/ustart_go/uses"

	getChat "github.com/sea350/ustart_go/get/chat"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	postChat "github.com/sea350/ustart_go/post/chat"
	post "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
)

//Page ... draws chat page
func Page(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	cs := client.ClientSide{}
	loadedMessages := []types.Message{}  //these variables made for legibility
	chatHeads := []types.FloatingHead{}  //not needed for functionality
	eaverHeads := []types.FloatingHead{} //

	chatID := r.URL.Path[4:]

	usr, err := get.UserByID(client.Eclient, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		cs.ErrorOutput = err
		cs.ErrorStatus = true
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderTemplate(w, r, "chat", cs)
		go handleMessages()
		return
	}

	if usr.ProxyMessagesID == `` {
		newProxy := types.ProxyMessages{DocID: docID.(string), Class: 1}
		proxyID, err := postChat.IndexProxyMsg(client.Eclient, newProxy)
		err = post.UpdateUser(client.Eclient, docID.(string), "ProxyMesssagesID", proxyID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			cs.ErrorOutput = err
			cs.ErrorStatus = true
			client.RenderSidebar(w, r, "template2-nil")
			client.RenderTemplate(w, r, "chat", cs)
			go handleMessages()
			return
		}
	} else {
		proxy, err := getChat.ProxyMsgByID(client.Eclient, usr.ProxyMessagesID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
			cs.ErrorOutput = err
			cs.ErrorStatus = true
			client.RenderSidebar(w, r, "template2-nil")
			client.RenderTemplate(w, r, "chat", cs)
			go handleMessages()
			return
		}

		for _, convoState := range proxy.Conversations {
			head, err := uses.ConvertChatToFloatingHead(client.Eclient, convoState.ConvoDocID, docID.(string))
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
				cs.ErrorOutput = errors.New("There were one or more errors loading chat heads")
				cs.ErrorStatus = true
			}
			chatHeads = append(chatHeads, head)
		}
	}

	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("debug text", len(chatID))
	// log.Println("debug text", chatID)

	if len(chatID) > 0 {
		if chatID[:1] == `@` {
			//this is a DM using username
			dmID, err := get.IDByUsername(client.Eclient, chatID[1:])
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
				http.Redirect(w, r, "/404/", http.StatusFound)
				return
			}

			exists, id, err := getChat.DMExists(client.Eclient, dmID, docID.(string))
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
				cs.ErrorOutput = err
				cs.ErrorStatus = true
				client.RenderSidebar(w, r, "template2-nil")
				client.RenderTemplate(w, r, "chat", cs)
				go handleMessages()
				return
			}
			if exists {
				convo, err := getChat.ConvoByID(client.Eclient, id)
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					dir, _ := os.Getwd()
					log.Println(dir, err)
					http.Redirect(w, r, "/404/", http.StatusFound)
					return
				}
				for _, eaver := range convo.Eavesdroppers {
					head, err := uses.ConvertUserToFloatingHead(client.Eclient, eaver.DocID)
					if err != nil {
						log.SetFlags(log.LstdFlags | log.Lshortfile)
						dir, _ := os.Getwd()
						log.Println(dir, err)
						cs.ErrorOutput = errors.New("There were one or more errors loading conversation members")
						cs.ErrorStatus = true
					} else {
						eaverHeads = append(eaverHeads, head)
					}
				}
				for _, msgid := range convo.MessageIDCache {
					msg, err := getChat.MsgByID(client.Eclient, msgid)
					if err != nil {
						log.SetFlags(log.LstdFlags | log.Lshortfile)
						dir, _ := os.Getwd()
						log.Println(dir, err)
						cs.ErrorOutput = errors.New("There were one or more errors loading conversation members")
						cs.ErrorStatus = true
					} else {
						loadedMessages = append(loadedMessages, msg)
					}
				}
			}
		} else {
			convo, err := getChat.ConvoByID(client.Eclient, chatID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
				http.Redirect(w, r, "/404/", http.StatusFound)
				return
			}
			_, exists := convo.Eavesdroppers[docID.(string)]
			if !exists {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, "THIS USER IS NOT PART OF THE CONVERSATION")
				http.Redirect(w, r, "/404/", http.StatusFound)
				return
			}
			for _, eaver := range convo.Eavesdroppers {
				head, err := uses.ConvertUserToFloatingHead(client.Eclient, eaver.DocID)
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					dir, _ := os.Getwd()
					log.Println(dir, err)
					cs.ErrorOutput = errors.New("There were one or more errors loading conversation members")
					cs.ErrorStatus = true
				}
				eaverHeads = append(eaverHeads, head)
			}
			for _, msgid := range convo.MessageIDCache {
				msg, err := getChat.MsgByID(client.Eclient, msgid)
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					dir, _ := os.Getwd()
					log.Println(dir, err)
					cs.ErrorOutput = errors.New("There were one or more errors loading conversation members")
					cs.ErrorStatus = true
				} else {
					loadedMessages = append(loadedMessages, msg)
				}
			}
		}
	}

	//get chat proxy
	//load list of heads
	cs.ListOfHeads = chatHeads
	cs.ListOfHeads2 = eaverHeads
	cs.Messages = loadedMessages
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderTemplate(w, r, "chat", cs)
	go handleMessages()
}
