package uses

import (
	"fmt"

	post "github.com/sea350/ustart_go/post/project"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeProjectLocation ... CHANGES PROJECT THE PROJECT'S LISTED LOCATION
//Requires the atarget project's docID all aspects of a types.LocStruct
//Returns an error if there was a problem with database submission
func ChangeProjectLocation(eclient *elastic.Client, projectID string, country string, state string, city string, zip string) error {
	var newLoc types.LocStruct
	newLoc.Country = country
	newLoc.State = state
	newLoc.City = city
	newLoc.Zip = zip

	fmt.Println("Country", country)
	fmt.Println("newloc", newLoc)
	err := post.UpdateProject(eclient, projectID, "Location", newLoc)
	return err

}
