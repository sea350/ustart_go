package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sea350/ustart_go/middleware/client"

	globals "github.com/sea350/ustart_go/globals"
)

var eclient = client.Eclient

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
