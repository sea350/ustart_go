package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"context"
	"encoding/json"
	"fmt"

	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	ctx := context.Background()

	maq := elastic.NewMatchAllQuery()
	res, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq).
		Size(100).
		Do(ctx)

	for _, id := range res.Hits.Hits {
		data := types.User{}
		err = json.Unmarshal(*id.Source, &data)
		if err != nil {
			fmt.Println(err)
		}
		if len(badgeIDs) == 0 {
			fmt.Println(data.FirstName, data.LastName, id.Id, data.Email)
			//
			// err = post.UpdateUser(eclient, id.Id, "BadgeIDs", append(data.BadgeIDs, badgeIDs...))

			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
