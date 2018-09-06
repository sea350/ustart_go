package post

import (
	"strings"

	elastic "gopkg.in/olivere/elastic.v5"
)

//ResetUnseen ... resets unseen to 0
//needs its own lock for concurrency control
func ResetUnseen(eclient *elastic.Client, proxyID string) error {

	ModifyUnseen.Lock()
	defer ModifyUnseen.Unlock()

	err := UpdateProxyNotifcations(eclient, strings.ToLower(proxyID), "NumUnseen", 0)

	return err
}
