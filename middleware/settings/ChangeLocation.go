package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeLocation ...  changes the user's geographical location in the database
func ChangeLocation(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	// userstruct, err := getUser.UserByID(client.Eclient, test1.(string))
	// countryInit := userstruct.Location.CountryVis
	// cityInit := userstruct.Location.CityVis
	// stateInit := userstruct.Location.StateVis
	// zipInit := userstruct.Location.ZipVis
	// if err != nil {
	// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 	dir, _ := os.Getwd()
	// 	log.Println(dir, err)
	// }
	r.ParseForm()
	countryP := r.FormValue("country")
	countryPV := r.FormValue("countryVis")
	//   fmt.Println(countryPV)
	stateP := r.FormValue("state")
	statePV := r.FormValue("stateVis")
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	// client.ClientSide{countryInit, cityInit, stateInit, zipInit}
	http.Redirect(w, r, "/Settings/#loccollapse", http.StatusFound)
	return

}
