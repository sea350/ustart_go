package post

import (
	elastic "github.com/olivere/elastic"
)

//ResetUnseen ... resets unseen to 0
//needs its own lock for concurrency control
func ResetUnseen(eclient *elastic.Client, proxyID string) error {

	ModifyUnseen.Lock()
	defer ModifyUnseen.Unlock()

	err := UpdateProxyNotifications(eclient, proxyID, "NumUnseen", 0)

	return err
}
