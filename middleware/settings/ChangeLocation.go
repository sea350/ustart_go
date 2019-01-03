package settings

import (
	"fmt"
	
	"net/http"
	

	"github.com/microcosm-cc/bluemonday"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeLocation ...  changes the user's geographical location in the database
func ChangeLocation(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// userstruct, err := getUser.UserByID(client.Eclient, test1.(string))
	// countryInit := userstruct.Location.CountryVis
	// cityInit := userstruct.Location.CityVis
	// stateInit := userstruct.Location.StateVis
	// zipInit := userstruct.Location.ZipVis
	// if err != nil {
	// 	
	// 	_ := os.Getwd()
	// 			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+err)
	// }
	r.ParseForm()
	p := bluemonday.UGCPolicy()
	countryP := r.FormValue("country")
	countryP = p.Sanitize(countryP)

	countryPV := r.FormValue("countryVis")
	countryPV = p.Sanitize(countryPV)
	//   fmt.Println(countryPV)
	stateP := r.FormValue("state")
	stateP = p.Sanitize(stateP)
	statePV := r.FormValue("stateVis")
	statePV = p.Sanitize(statePV)
	cityP := r.FormValue("city")
	cityPV := r.FormValue("cityVis")
	zipP := r.FormValue("zip")
	zipPV := r.FormValue("zipVis")
	conBool := false
	if countryPV == "on" {
		conBool = true
	}
	sBool := false
	if statePV == "on" {
		sBool = true
	}
	cBool := false
	if cityPV == "on" {
		cBool = true
	}
	zBool := false
	if zipPV == "on" {
		zBool = true
	}

	err := uses.ChangeLocation(client.Eclient, session.Values["DocID"].(string), countryP, conBool, stateP, sBool, cityP, cBool, zipP, zBool)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	// client.ClientSide{countryInit, cityInit, stateInit, zipInit}
	http.Redirect(w, r, "/Settings/#loccollapse", http.StatusFound)
	return

}
