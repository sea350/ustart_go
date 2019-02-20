//class 4 for follow, refID L, docID R
package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"context"
	"encoding/json"
	"fmt"

	// postUser "github.com/sea350/ustart_go/post/user"
	// postNotif "github.com/sea350/ustart_go/post/notification"
	getUser "github.com/sea350/ustart_go/get/user"
	// getNotif "github.com/sea350/ustart_go/get/notification"
	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	ctx := context.Background()

	LiD, _ := getUser.IDByUsername(eclient, "ln961")
	RiD, _ := getUser.IDByUsername(eclient, "ryanrozbiani")

	fmt.Println("IDs:", LiD, RiD)
	maq := elastic.NewBoolQuery()
	maq = maq.Must(elastic.NewTermQuery("DocID", RiD))
	maq = maq.Must(elastic.NewTermQuery("ReferenceIDs", LiD))
	res, err := eclient.Search().
		Index(globals.NotificationIndex).
		Type(globals.NotificationType).
		Query(maq).
		Size(100).
		Do(ctx)

	fmt.Println("nHits:", res.TotalHits())
	for _, id := range res.Hits.Hits {
		data := types.Notification{}
		err = json.Unmarshal(*id.Source, &data)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("TIMESTAMP:", data.Timestamp)
		//
		// err = post.UpdateUser(eclient, id.Id, "BadgeIDs", append(data.BadgeIDs, badgeIDs...))

		if err != nil {
			fmt.Println(err)
			continue

		}
	}
}
