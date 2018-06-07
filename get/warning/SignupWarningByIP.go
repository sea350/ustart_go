package get

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_go/types"
)

//SingupWarningByIP ...
//Retrive Signup Warning structure based off of the addressIP of the user
func SingupWarningByIP(eclient *elastic.Client, addressIP string) (types.SignupWarning, error) {
	ctx := context.Background()
	var signWarning types.SignupWarning
	termQuery := elastic.NewTermQuery("IPAddress", addressIP)
	searchResult, err := eclient.Search().Index().Query(termQuery).Do(ctx)
	if err != nil {
		return signWarning, err
	}

	if searchResult.Hits.TotalHits == 0 {
		err1 := AppendIndex(eclient, signWarning)
		return signWarning, err1
	} else {
		var ipID string
		for _, res := range searchResult.Hits.Hits {
			ipID = res.Id
			break
		}

		initSignWarning, err2 := eclient.Get().Index("ipIndex").Id(ipID).Do(ctx)
		if err2 != nil {
			panic(err2)
		}

		err3 := json.Unmarshal(*initSignWarning.Source, &signWarning)
		return signWarning, err3
	}
}
