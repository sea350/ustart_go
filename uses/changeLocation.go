package uses

import (
	post "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ChangeLocation ...
func ChangeLocation(eclient *elastic.Client, userID string, country string, countryVis bool, state string, stateVis bool, city string, cityVis bool, zip string, zipVis bool) error {

	var newLoc types.LocStruct
	newLoc.Country = country
	newLoc.CountryVis = countryVis
	newLoc.State = state
	newLoc.StateVis = stateVis
	newLoc.City = city
	newLoc.CityVis = cityVis
	newLoc.Zip = zip
	newLoc.ZipVis = zipVis
	err := post.UpdateUser(eclient, userID, "Location", newLoc)
	return err

}
