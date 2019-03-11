package main

import (
	"fmt"

	getCode "github.com/sea350/ustart_go/get/guestCode"
	"github.com/sea350/ustart_go/middleware/client"
)

func main() {

	//initialize bool query
	// boolSearch := elastic.NewBoolQuery()
	// searchResults, err := client.Eclient.Search().
	// 	Index(globals.GuestCodeIndex).
	// 	Query(boolSearch).
	// 	Do(context.Background())

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// code := types.GuestCode{}

	// for _, element := range searchResults.Hits.Hits {
	// 	fmt.Println(element.Id)
	// 	err := json.Unmarshal(*element.Source, &code) //unmarshal type RawMessage into user struct
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	fmt.Println(code.NumUses)
	// 	fmt.Println(code.Expiration)
	// }

	code, err := getCode.GuestCodeByID(client.Eclient, "NYUFEST")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(code.Users)

}
