package post

import (
	"context"
	"errors"
	"fmt"

	get "github.com/sea350/ustart_go/get/warning"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateSignUpWarningByIP ...
//A single field update (not used)
func UpdateSignUpWarningByIP(eclient *elastic.Client, addressIP string, field string, newContent interface{}) error {
	//code
	ctx := context.Background()
	exists, err := eclient.IndexExists("ipindex").Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.SingupWarningByIP(eclient, addressIP)
	if err != nil {
		return err
	}

	termQuery := elastic.NewTermQuery("IPAddress", addressIP)
	searchResult, err := eclient.Search().Index("ipindex").Query(termQuery).Do(ctx)
	var ipID string
	for _, res := range searchResult.Hits.Hits {
		ipID = res.Id
		break
	}

	_, err = eclient.Update().
		Index("ipindex").
		Type("IPADDRESS").
		Id(ipID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	return err

}

//ReIndexSingupWarning ...
//Updates ALL fields
func ReIndexSignupWarning(eclient *elastic.Client, signWarning types.SignupWarning, addressIP string) error {
	ctx := context.Background()
	exists, err := eclient.IndexExists("ipindex").Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.SingupWarningByIP(eclient, addressIP)
	if err != nil {
		return err
	}

	//Fix Query
	matchQuery := elastic.NewMatchQuery("IPAddress", addressIP) //From NewTermQuery
	searchResult, err := eclient.Search().Index("ipindex").Query(matchQuery).Do(ctx)
	var ipID string
	for _, res := range searchResult.Hits.Hits {
		ipID = res.Id
		break
	}

	fmt.Println("Here is the ID", ipID)

	_, err = eclient.Index().
		Index("ipindex").
		Type("IPADDRESS").
		Id(ipID).
		BodyJson(signWarning).
		Do(ctx)

	return err
}
