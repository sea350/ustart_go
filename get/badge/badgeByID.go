package get

import (
	"context"
	"encoding/json"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
	//post "github.com/sea350/ustart_go/post"
)

//BadgeByID ...
func BadgeByID(eclient *elastic.Client, badgeID string) (types.Badge, error) {
	ctx := context.Background()         //intialize context background
	var badge types.Badge               //initialize type user
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.BadgeIndex).
						Type(globals.BadgeType).
						Id(badgeID).
						Do(ctx)

	if err != nil {
		return badge, err
	}

	Err := json.Unmarshal(*searchResult.Source, &badge) //unmarshal type RawMessage into user struct

	return badge, Err

}
