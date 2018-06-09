package get

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sea350/ustart_go/types"
	"gopkg.in/olivere/elastic.v5"
)

const ipMapping = `
	"mappings":{
		"IPADDRESS":{
			"properties":{	
				"IPAddress":{
					"type":"keyword"
					}
					
				}
			}
		}
	}
`

func startIndex(eclient *elastic.Client) error {
	ctx := context.Background()

	_, err := eclient.CreateIndex("ipindex").BodyString(ipMapping).Do(ctx)
	if err != nil {
		fmt.Println("Could not create", ipMapping)
		return err

	} else {
		fmt.Println(ipMapping, "created")
	}
	return err

}

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

	exists, err := eclient.IndexExists("ipindex").Do(ctx)
	if err != nil {
		return signWarning, err
	}
	if !exists {
		err := startIndex(eclient)
		if err != nil {
			return signWarning, err
		}
	}

	if searchResult.Hits.TotalHits == 0 {
		err1 := AppendIndexSignWarning(eclient, signWarning)
		return signWarning, err1
	} else {
		var ipID string
		for _, res := range searchResult.Hits.Hits {
			ipID = res.Id
			break
		}

		initSignWarning, err2 := eclient.Get().Index("ipindex").Id(ipID).Do(ctx)
		if err2 != nil {
			panic(err2)
		}

		err3 := json.Unmarshal(*initSignWarning.Source, &signWarning)
		return signWarning, err3
	}
}
