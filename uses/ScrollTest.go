package uses

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	elastic "github.com/olivere/elastic"
)

//ScrollTest ...
//Trying to understand how the elastic scroll service works (attempting to be used with LoadEntires.go function)
func ScrollTest(eclient *elastic.Client, loadlist []string) {
	/* Maybe this?
	var newString string
	for i := 0; i < len(loadlist); i++ {
		newString += loadlist[i]
	}
	sliceQuery := elastic.NewSliceQuery().Field(newString).Id(0).Max(len(loadlist))
	svc := eclient.Scroll().Slice(sliceQuery)
	*/

	sliceQuery := elastic.NewSliceQuery().Id(0).Max(len(loadlist))
	svc := eclient.Scroll("TestIndex").Slice(sliceQuery)
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
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		if res == nil {
			fmt.Println("expected results != nil; got nil")
		}
		if res.Hits == nil {
			fmt.Println("expected results.Hits != nil; got nil")
		}

		pages++

		for _, hit := range res.Hits.Hits {
			if hit.Index != "TestIndex" {
				fmt.Println("expected SearchResult.Hits.Hit.Index = TestIndex; got " + hit.Index)
			}
			item := make(map[string]interface{})
			err := json.Unmarshal(*hit.Source, &item)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
			}
			docs++
		}

		if len(res.ScrollId) == 0 {
			fmt.Println("expected scrollId in results; got " + res.ScrollId)
		}

	}

	if pages == 0 {
		fmt.Println("expected to retrieve some pages")
	}
	if docs == 0 {
		fmt.Println("expected to retrieve some hits")
	}

	if err := svc.Clear(context.TODO()); err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	if _, err := svc.Do(context.TODO()); err == nil {
		fmt.Println("expected to fail")
	}
}
