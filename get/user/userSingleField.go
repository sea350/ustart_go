package get

import (
	"context"
	"encoding/json"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//SingleStringField ...
//Retreives a single field from user
func SingleStringField(eclient *elastic.Client, queryField string, queryTerm string, includeField string) (string, error) {

	ctx := context.Background()

	//username:= EmailToUsername(email) //for username query
	termQuery := elastic.NewTermQuery("Username", "jc5537")

	fsc := elastic.NewFetchSourceContext(true).Include("Tags") //.Exclude("*")
	builder := elastic.NewSearchSource().Query(termQuery).DocvalueField("hello").FetchSourceContext(fsc)
	// src, err := builder.Source()
	searchRes, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		SearchSource(builder).
		Do(ctx)

	var data []byte
	// Err := json.Unmarshal(*searchRes.Source, &usr)
	for _, element := range searchRes.Hits.Hits {

		result := element.Source
		data, _ = json.Marshal(*result)
		break

	}

	// fmt.Println()
	fmt.Println(string(data))
	// var usr types.User
	if err != nil {
		fmt.Println(err)

	}
	return string(data), err

}