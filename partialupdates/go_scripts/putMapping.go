package partialupdates

import (
	"context"
	"fmt"

	elastic "github.com/olivere/elastic"
)

//PutMapping ...
//Modifies existing mappings
func PutMapping(eclient *elastic.Client, index string, docType string, mapping string) error {
	ctx := context.Background()
	res, err := eclient.PutMapping().
		Index(index).
		Type(docType).
		BodyString(mapping).
		Do(ctx)

	fmt.Println("DEBUG IN PUTMAPPING:", res)
	return err
}
