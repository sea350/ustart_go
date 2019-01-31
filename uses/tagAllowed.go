package uses

import (
	"context"
	"log"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
	//post "github.com/sea350/ustart_go/post"
)

//ByID ...
func TagAllowed(eclient *elastic.Client, newTag string) bool {
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

	return res.Hits.TotalHits == 0

}
