package get

import (
	"context"

	"github.com/sea350/ustart_go/types"
	"gopkg.in/olivere/elastic.v5"
)

//AppendIndexSignWarning ...
func AppendIndexSignWarning(eclient *elastic.Client, signUpWarning types.SignupWarning) error {
	ctx := context.Background()
	_, err := eclient.Index().Index("ipIndex").Type("IPADDRESS").BodyJson(signUpWarning).Do(ctx)
	return err

}
