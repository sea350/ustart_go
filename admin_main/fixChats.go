package main

import (
	"context"
	"fmt"

	getChat "github.com/sea350/ustart_go/get/chat"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//9v4r-GgBN3VvtvdieZzG
// g_5h42gBN3VvtvdiWZt3

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {
	ctx := context.Background()
	// proxyID := "g_5h42gBN3VvtvdiWZt3"
	// usrID, _ := getUser.IDByUsername(eclient, "HeatherMT")
	convoID := "9v4r-GgBN3VvtvdieZzG"

	// fmt.Println("Deleting convo")
	// _, err := eclient.Delete().
	// 	Index(globals.ConvoIndex).
	// 	Type(globals.ConvoType).
	// 	Id(focusID).
	// 	Do(ctx)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var theProxy = types.ProxyMessages{
	// 	DocID:         usrID,
	// 	Class:         1,
	// 	NumUnread:     0,
	// 	Conversations: nil,
	// }

	// _, err := eclient.Index().
	// 	Index(globals.ProxyMsgIndex).
	// 	Type(globals.ProxyMsgType).
	// 	Id(proxyID).
	// 	BodyJson(theProxy).
	// 	Do(ctx)

	// var blankTime time.Time
	// var convoState = types.ConversationState{
	// 	// NumUnread   int       `json:"NumUnread"`
	// 	// LastMessage Message   `json:"LastMessage"`
	// 	ConvoID:     "-f6_6WgBN3VvtvdiTJtI",
	// 	ProjectID:   "9_6_6WgBN3VvtvdiTJsk",
	// 	Read:        true,
	// 	Muted:       false,
	// 	MuteTimeout: blankTime,
	// }

	// var convoState2 = types.ConversationState{
	// 	// NumUnread   int       `json:"NumUnread"`
	// 	// LastMessage Message   `json:"LastMessage"`
	// 	ConvoID:     "9P4r-GgBN3Vvtvdicpzp",
	// 	ProjectID:   "",
	// 	Read:        true,
	// 	Muted:       false,
	// 	MuteTimeout: blankTime,
	// }

	// convoStates := []types.ConversationState{convoState, convoState2}

	// proxyObj, err := getChat.ProxyMsgByID(eclient, proxyID)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = postChat.UpdateProxyMsg(eclient, proxyID, "Conversations", append(proxyObj.Conversations, convoStates...))

	// if err != nil {
	// 	fmt.Println(err)
	// }

	ntq := elastic.NewTermsQuery("_id", convoID)

	res, err := eclient.Search().
		Index(globals.ConvoIndex).
		Query(ntq).
		Do(ctx)

	for _, r := range res.Hits.Hits {
		cvo, _ := getChat.ConvoByID(eclient, r.Id)
		fmt.Println(cvo)
	}
}
