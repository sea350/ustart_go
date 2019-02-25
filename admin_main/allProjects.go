package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	elastic "github.com/olivere/elastic"
	// getUser "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/globals"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

//Jv63yWgBN3Vvtvdiu5YP

func main() {

	ctx := context.Background()

	maq := elastic.NewMatchAllQuery()
	res, err := eclient.Search().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Query(maq).
		Size(500).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, id := range res.Hits.Hits {
		data := types.User{}
		err = json.Unmarshal(*id.Source, &data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data.Name "  ", data.URLName)
	}

}
