package main

import (
	"context"
	"fmt"
	"log"

	// admin "github.com/sea350/ustart_go/admin"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func ta(eclient *elastic.Client, newTag string) bool {
	ctx := context.Background()

	tagQuery := elastic.NewTermQuery("Tags", newTag)

	res, err := eclient.Search().
		Index(globals.BadgeIndex).
		Type(globals.BadgeType).
		Query(tagQuery).
		Do(ctx)

	if err != nil {
		log.Println(err)
		return false
	}

	// var badgeID string
	// for _, hit := range res.Hits.Hits {
	// 	badgeID = hit.Id
	// 	break
	// }

	// userBadgeQuery := elastic.NewBoolQuery()
	// userBadgeQuery = userBadgeQuery.Must(elastic.NewTermQuery("BadgeIDs", badgeID)).Must(elastic.NewTermQuery("_id", usrID))

	// badgeRes, err := eclient.Search().
	// 	Index(globals.BadgeIndex).
	// 	Type(globals.BadgeType).
	// 	Query(userBadgeQuery).
	// 	Do(ctx)

	// if err != nil {
	// 	log.Println(err)
	// 	return false
	// }

	return res.Hits.TotalHits == 0 //&& badgeRes.Hits.TotalHits == 1

}

func main() {

	fmt.Println(ta(eclient, "U·START VIP Spring 2019"))
}
