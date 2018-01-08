package uses

import (
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	post "github.com/sea350/ustart_go/post/project"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeProjectURL ... CHANGES PROJECT THE PROJECT'S URL EXTENTION
//Requires the target projects docID and the potential new url
//Returns an error if the url is taken or a databse error
func ChangeProjectURL(eclient *elastic.Client, projectID string, newURL string) error {
	_, err := get.ProjectByURL(eclient, newURL)
	//if (err != nil){ return err}
	if err != nil {
		return errors.New("That url is already taken")
	}
	err = post.UpdateProject(eclient, projectID, "URLName", newURL)
	return err
}
