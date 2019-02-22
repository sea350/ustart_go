package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"net/http"

	get "github.com/sea350/ustart_go/get/notification"
	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/notification"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "github.com/olivere/elastic"
)

//AjaxNotificationLoad ... crawling in the 90s
//Designed for ajax
func AjaxNotificationLoad(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	var notifs []map[string]interface{}

	id, proxy, err := get.ProxyNotificationByUserID(client.Eclient, docID.(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	err = post.ResetUnseen(client.Eclient, id)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	// count := 0
	// for _, id := range proxy.NotificationCache {
	// 	notif, err := get.NotificationByID(client.Eclient, id)
	// 	if err != nil {

	// 		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	// 		continue
	// 	}
	// 	if notif.Invisible {
	// 		continue
	// 	}

	// 	msg, url, err := uses.GenerateNotifMsgAndLink(client.Eclient, notif)
	// 	if err != nil {

	// 		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	// 		continue
	// 	}

	// 	notifAggregate := make(map[string]interface{})
	// 	notifAggregate["ID"] = id
	// 	notifAggregate["Data"] = notif
	// 	notifAggregate["Message"] = msg
	// 	notifAggregate["URL"] = url
	// 	notifs = append(notifs, notifAggregate)
	// 	count++
	// 	if count == 5 {
	// 		break
	// 	}

	// }

	ctx := context.Background()

	//set up user query
	notifQuery := elastic.NewBoolQuery()
	notifQuery = notifQuery.Must(elastic.NewTermQuery("DocID", strings.ToLower(docID.(string))))
	notifQuery = notifQuery.Must(elastic.NewTermQuery("Invisible", false))

	inboxQuery := client.Eclient.Search().
		Index(globals.NotificationIndex).
		Query(notifQuery).
		Sort("Timestamp", false).
		Size(5)

	res, err := inboxQuery.Do(ctx)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		return
	}

	var notif types.Notification
	for _, hit := range res.Hits.Hits {
		// fmt.Println(hit.Id)
		if err == io.EOF {
			continue
		}

		err := json.Unmarshal(*hit.Source, &notif)
		if err != nil && err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}

		msg, url, err := uses.GenerateNotifMsgAndLink(client.Eclient, notif)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			continue
		}

		notifAggregate := make(map[string]interface{})
		notifAggregate["ID"] = hit.Id
		notifAggregate["Data"] = notif
		notifAggregate["Message"] = msg
		notifAggregate["URL"] = url
		notifs = append(notifs, notifAggregate)
	}

	sendData := make(map[string]interface{})
	sendData["notifications"] = notifs

	//set up user query
	// unreadQuery := elastic.NewBoolQuery()
	// unreadQuery = unreadQuery.Must(elastic.NewTermQuery("DocID", strings.ToLower(docID.(string))))
	// unreadQuery = unreadQuery.Must(elastic.NewTermQuery("Invisible", false))
	// unreadQuery = unreadQuery.Must(elastic.NewTermQuery("Seen", false))

	// res, err = client.Eclient.Search().
	// 	Index(globals.NotificationIndex).
	// 	Query(notifQuery).
	// 	Sort("Timestamp", false).
	// 	Do(ctx)

	// if err != nil {
	// 	client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	// }
	sendData["numUnread"] = proxy.NumUnread

	data, err := json.Marshal(sendData)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	fmt.Fprintln(w, string(data))
}
