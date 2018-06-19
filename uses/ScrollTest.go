package uses

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	elastic "gopkg.in/olivere/elastic.v5"
)

//ScrollTest ...
//Trying to understand how the elastic scroll service works (attempting to be used with LoadEntires.go function)
func ScrollTest(eclient *elastic.Client, loadlist []string) {

	//1st version
	var newString string
	for i := 0; i < len(loadlist); i++ {
		newString += loadlist[i]
	}
	sliceQuery := elastic.NewSliceQuery().Field(newString).Id(0).Max(len(loadlist))
	//EntryIDs for now, but maybe can fill in for Project EntryIDS or anything else?
	svc := eclient.Scroll("EntryIDs").Slice(sliceQuery)

	//2nd version
	/*sliceQuery := elastic.NewSliceQuery().Id(0).Max(len(loadlist))
	svc := eclient.Scroll("TestIndex").Slice(sliceQuery)
	for i := 0; i < len(loadlist); i++ {
		svc.Index(loadlist[i])
	}
	*/

	//Pulled from https://github.com/olivere/elastic/blob/release-branch.v6/search_queries_slice_test.go
	//Slice Test

	/*
		src, err := sliceQuery.Source()
		if err != nil {
			fmt.Println(err)
		}
		data, err := json.Marshal(src)
		if err != nil {
			fmt.Println("marshaling to JSON failed: %v", err)
		}
		got := string(data)
		if got != expected {
			fmt.Println("got: ", got)
		}
	*/
	//Pulled from https://github.com/olivere/elastic/blob/release-branch.v6/scroll_test.go
	//Scroll Test

	pages := 0
	docs := 0

	for {
		res, err := svc.Do(context.TODO())
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if res == nil {
			fmt.Println("expected results != nil; got nil")
		}
		if res.Hits == nil {
			fmt.Println("expected results.Hits != nil; got nil")
		}

		pages++

		for _, hit := range res.Hits.Hits {
			if hit.Index != "EntryIDs" {
				fmt.Println("Expected EntryIDs index, got: ", hit.Index)
			}
			item := make(map[string]interface{})
			err := json.Unmarshal(*hit.Source, &item)
			if err != nil {
				fmt.Println(err)
			}
			docs++
		}

		if len(res.ScrollId) == 0 {
			fmt.Println("expected scrollId in results; got: ", res.ScrollId)
		}

	}

	if pages == 0 {
		fmt.Println("expected to retrieve some pages")
	}
	if docs == 0 {
		fmt.Println("expected to retrieve some hits")
	}

	if err := svc.Clear(context.TODO()); err != nil {
		fmt.Println(err)
	}

	if _, err := svc.Do(context.TODO()); err == nil {
		fmt.Println("expected to fail")
	}

}
