package uses

import (
	"errors"

	get "github.com/sea350/ustart_go/get/event"
	post "github.com/sea350/ustart_go/post/event"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeEventURL ... CHANGES EVENT THE EVENT'S URL EXTENTION
//Requires the target events docID and the potential new url
//Returns an error if the url is taken or a databse error
func ChangeEventURL(eclient *elastic.Client, eventID string, newURL string) error {
	inUse, err := get.EventURLInUse(eclient, newURL)
	//if (err != nil){ return err}
	if inUse {
		return errors.New("That url is already taken")
	}
	err = post.UpdateEvent(eclient, eventID, "URLName", newURL)
	return err
}
