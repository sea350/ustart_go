package post

import (
	"context"

	"github.com/sea350/ustart_go/types"
	"github.com/olivere/elastic"
)

//AddToIndexSignWarning ...
func AddToIndexSignWarning(eclient *elastic.Client, signUpWarning types.SignupWarning) error {
	ctx := context.Background()
	_, err := eclient.Index().Index("ipindex").Type("IPADDRESS").BodyJson(signUpWarning).Do(ctx)
	return err

}
