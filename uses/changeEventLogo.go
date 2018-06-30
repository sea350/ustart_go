package uses

import (
	post "github.com/sea350/ustart_go/post/event"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeEventLogo ...
func ChangeEventLogo(eclient *elastic.Client, eventID string, newLogo string) error {

	return post.UpdateEvent(eclient, eventID, "Avatar", newLogo)

}
