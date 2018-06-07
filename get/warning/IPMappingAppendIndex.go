package get

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_go/types"
)

//AppendIndex ...
func AppendIndex(eclient *elastic.Client, signUpWarning types.SignupWarning) error {
	ctx := context.Background()
	_, err := eclient.Index().Index("ipIndex").Type("IPADDRESS").BodyJson(signUpWarning).Do(ctx)
	return err

}
