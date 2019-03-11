package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
)

func main() {

	//initialize bool query
	boolSearch := elastic.NewBoolQuery()
	searchResults, err := eclient.Search().
		Index(globals.GuestCodeIndex).
		Query(boolSearch).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	code := types.GuestCode{}

	for _, element := range searchResults.Hits.Hits {
		fmt.Println(element.Id)
		err := json.Unmarshal(*element.Source, &code) //unmarshal type RawMessage into user struct
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(code.NumUses)
	}

}
