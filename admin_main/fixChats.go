package main

import (
	"context"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//9v4r-GgBN3VvtvdieZzG
// g_5h42gBN3VvtvdiWZt3

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {
	ctx := context.Background()
	proxyID := " g_5h42gBN3VvtvdiWZt3"
	focusID := "9v4r-GgBN3VvtvdieZzG"

	fmt.Println("Deleting convo")
	err := eclient.Delete().
		Index(globals.ConvoIndex).
		Type(globals.ConvoType).
		Id(focusID).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Deleting proxy")
	err = eclient.Delete().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		Id(proxyID).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

}
