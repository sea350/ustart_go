package post

import (
	"context"
	"errors"

	elastic "github.com/olivere/elastic"
)

//UpdateSignUpWarningByIP ...
func UpdateSignUpWarningByIP(eclient *elastic.Client, addressIP string, field string, newContent interface{}) error {
	//code
	ctx := context.Background()
	exists, err := eclient.IndexExists("ipIndex").Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	//signWarning, err := get.SingupWarningByIP(eclient, addressIP)
	//if err != nil {
	//	return err
	//}

	//_, err := eclient.Update().Index("ipIndex").Id(ipID).Do(ctx)
	return err

}
