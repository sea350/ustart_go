package get

import (
	"context"
	"encoding/json"
	"log"

	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//SingleStringField ...
//Retreives a single field from user
func SingleStringField(eclient *elastic.Client, queryField string, queryTerm string, includeField string) (string, error) {

	ctx := context.Background()

	//username:= EmailToUsername(email) //for username query
	termQuery := elastic.NewTermQuery(queryField, queryTerm)

	fsc := elastic.NewFetchSourceContext(true).Include(includeField) //.Exclude("*")
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

	// var usr types.User
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)

	}
	return string(data), err

}

// //StringArrayField ...
// //Retreives a single string array field from user
// func StringArrayField(eclient *elastic.Client, queryField string, queryTerm string, includeField string) ([]string, error) {

// 	ctx := context.Background()

// 	//username:= EmailToUsername(email) //for username query
// 	termQuery := elastic.NewTermQuery(queryField, queryTerm)

// 	fsc := elastic.NewFetchSourceContext(true).Include(includeField) //.Exclude("*")
// 	builder := elastic.NewSearchSource().Query(termQuery).DocvalueField("hello").FetchSourceContext(fsc)
// 	// src, err := builder.Source()
// 	searchRes, err := eclient.Search().
// 		Index(globals.UserIndex).
// 		Type(globals.UserType).
// 		SearchSource(builder).
// 		Do(ctx)

// 	var data []byte
// 	// Err := json.Unmarshal(*searchRes.Source, &usr)
// 	for _, element := range searchRes.Hits.Hits {

// 		result := element.Source
// 		data, _ = json.Marshal(*result)
// 		break

// 	}

// 	// var usr types.User
// 	if err != nil {
// 		log.SetFlags(log.LstdFlags | log.Lshortfile)
//		log.Println(err)

// 	}
// 	return []string(data), err

// }
