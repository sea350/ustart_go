package main

import (
	"context"
	"fmt"
	"log"
	"os"

	globals "github.com/sea350/ustart_go/backend/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL("http://localhost:9200"))

//DeleteAll ... nukes all indexes in ES
func DeleteAll() {

	ctx := context.Background()
	deleteIndex, err := eclient.DeleteIndex(globals.ChatIndex).Do(ctx)
	deleteIndex, err = eclient.DeleteIndex(globals.EntryIndex).Do(ctx)
	deleteIndex, err = eclient.DeleteIndex(globals.ProjectIndex).Do(ctx)
	deleteIndex, err = eclient.DeleteIndex(globals.WidgetIndex).Do(ctx)
	deleteIndex, err = eclient.DeleteIndex(globals.UserIndex).Do(ctx)
	if err != nil {
		// Handle error
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	} else {
		fmt.Println("S U C C")
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
	}
}
