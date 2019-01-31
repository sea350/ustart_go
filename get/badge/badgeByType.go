package get

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/badge"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//BadgeByType ...
//REtreives badge by type
func BadgeByType(eclient *elastic.Client, badgeType string) (types.Badge, error) {
	ctx := context.Background() //intialize context background

	bQuery := elastic.NewTermQuery("Type", badgeType)
	searchResult, err := eclient.Search(). //Get returns doc type, index, etc.
						Index(globals.BadgeIndex).
						Type(globals.BadgeType).
						Query(bQuery).
						Do(ctx)

	if err != nil {
		return false, err
	}

	var result string
	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		break
	}

	if len(result) == 0 {
		return badge, errors.New("ID Not found")
	}
	badge, err = get.BadgeByID(eclient, result)

	return badge, err

}
