package uses

import (
	"context"
	"fmt"
	"io"

	elastic "gopkg.in/olivere/elastic.v5"
)

//ScrollTest ...
func ScrollTest(eclient *elastic.Client, loadlist []string) {
	/*
		var newString string
		for i := 0; i < len(loadlist); i++ {
			newString += loadlist[i]
		}
		sliceQuery := elastic.NewSliceQuery().Field(newString).Id(0).Max(len(loadlist))
		svc := eclient.Scroll().Slice(sliceQuery)
	*/

	sliceQuery := elastic.NewSliceQuery().Id(0).Max(len(loadlist))
	svc := eclient.Scroll().Slice(sliceQuery)
	for i := 0; i < len(loadlist); i++ {
		svc.Index(loadlist[i])
	}

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
		/*
			for _, hit := range res.Hits.Hits {
				if hit.Index != testIndexName {
					fmt.Println("expected SearchResult.Hits.Hit.Index = %q; got %q", testIndexName, hit.Index)
				}
				item := make(map[string]interface{})
				err := json.Unmarshal(*hit.Source, &item)
				if err != nil {
					fmt.Println(err)
				}
				docs++
			}

			if len(res.ScrollId) == 0 {
				fmt.Println("expected scrollId in results; got %q", res.ScrollId)
			}
		*/
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
