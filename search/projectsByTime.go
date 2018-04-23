package search

import (
	"context"
	"encoding/json"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//ProjectsByTime ...
//Searching projects within a specific time range
func ProjectsByTime(eclient *elastic.Client, minTime string, maxTime string) {
	ctx := context.Background()
	queryThis := "Nil"
	idx := globals.ProjectIndex
	typ := globals.ProjectType

	bq := elastic.NewBoolQuery()
	bq = bq.Must(elastic.NewTermQuery("FirstName", queryThis))
	bq = bq.Must(elastic.NewRangeQuery("AccCreation").From("2017-01-01").To("2018-04-19").Boost(3))
	q := elastic.NewNestedQuery("Project", bq).QueryName("qname")
	src, err := q.Source()

	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(src)
	if err != nil {
		//t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)

	if got == got {
	}

	res, _ := eclient.Search().
		Index(idx).
		Type(typ).
		Query(bq).
		Do(ctx)
	//fmt.Println(res)
	//fmt.Println(res.Suggest)

	//fmt.Println(got)
	fmt.Println(res.Hits.TotalHits)
	fmt.Println(res.TookInMillis)

}