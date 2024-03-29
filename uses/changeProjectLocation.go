package uses

import (
	"fmt"

	post "github.com/sea350/ustart_go/post/project"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ChangeProjectLocation ... CHANGES PROJECT THE PROJECT'S LISTED LOCATION
//Requires the atarget project's docID all aspects of a types.LocStruct
//Returns an error if there was a problem with database submission
func ChangeProjectLocation(eclient *elastic.Client, projectID string, country string, state string, city string, zip string, vis bool) error {
	var newLoc types.LocStruct
	newLoc.Country = country
	newLoc.State = state
	newLoc.City = city
	newLoc.Zip = zip
	newLoc.CountryVis = vis
	newLoc.StateVis = vis
	newLoc.CityVis = vis
	newLoc.ZipVis = vis

	fmt.Println("Country", country)
	fmt.Println("newloc", newLoc)
	err := post.UpdateProject(eclient, projectID, "Location", newLoc)
	return err

}
