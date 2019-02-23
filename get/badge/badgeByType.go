package get

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//BadgeByType ...
//REtreives badge by type
func BadgeByType(eclient *elastic.Client, badgeType string) (types.Badge, error) {
	ctx := context.Background() //intialize context background

	bQuery := elastic.NewTermQuery("Type", strings.ToLower(badgeType))
	searchResult, err := eclient.Search(). //Get returns doc type, index, etc.
						Index(globals.BadgeIndex).
						Type(globals.BadgeType).
						Query(bQuery).
						Do(ctx)

	var badge types.Badge
	if err != nil {
		return badge, err
	}

	var result string
	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		break
	}

	if len(result) == 0 {
		return badge, errors.New("ID Not found")
	}
	badge, err = BadgeByID(eclient, result)

	return badge, err

}
