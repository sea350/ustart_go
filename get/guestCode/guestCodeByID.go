package get

import (
	"context"
	"encoding/json"

	"github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//GuestCodeByID ... Gets a guestCode by it's ID
func GuestCodeByID(eclient *elastic.Client, codeID string) (types.GuestCode, error) {
	ctx := context.Background()
	var guestCode types.GuestCode

	searchResults, err := eclient.Get().
		Index(globals.GuestCodeIndex).
		Id(codeID).
		Do(ctx)
	if err != nil {
		return guestCode, err
	}

	err = json.Unmarshal(*searchResults.Source, &guestCode)
	if err != nil {
		return guestCode, err
	}

	return guestCode, err
}
