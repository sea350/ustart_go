package search

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/sea350/ustart_go/backend/middleware/client"
	"github.com/sea350/ustart_go/backend/search"
	"github.com/sea350/ustart_go/backend/types"
	"github.com/sea350/ustart_go/backend/uses"
)

//AjaxLoadNext ... Loads more results after initial results loads once bottom of page is reached
//Takes w, r and search parameters from form values
//Returns Marshalled results(array of floating heads)
func AjaxLoadNext(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	var searchBy []bool
	query := r.FormValue("query")
	scrollID := r.FormValue("scrollID")
	filter := r.FormValue("searchFilterGroup") //can be: skills,users,projects
	searchMajors := uses.ConvertStrToStrArr(r.FormValue("searchlistmajors"))
	searchSkills := uses.ConvertStrToStrArr(r.FormValue("searchlistskills")) //array

	if filter == `projects` {
		if r.FormValue("searchbyprojectname") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbyurl") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbymembersneeded") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbyskills") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		totalHits, scrollID, results, err := search.PrototypeProjectSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, scrollID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		sendThis := make(map[string]interface{})
		sendThis["TotalHits"] = totalHits
		sendThis["ScrollID"] = scrollID
		sendThis["Results"] = results
		data, err := json.Marshal(sendThis)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		fmt.Fprintln(w, string(data))
	}
	if filter == `events` {
		if r.FormValue("searchbyeventname") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbyurl") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbymembers") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbyguests") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbyskills") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		totalHits, scrollID, results, err := search.PrototypeEventSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, scrollID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		sendThis := make(map[string]interface{})
		sendThis["TotalHits"] = totalHits
		sendThis["ScrollID"] = scrollID
		sendThis["Results"] = results
		data, err := json.Marshal(sendThis)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		fmt.Fprintln(w, string(data))
	}
	if filter == `users` {
		if r.FormValue("searchbypersonname") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbyusername") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		if r.FormValue("searchbyskills") != `` {
			searchBy = append(searchBy, true)
		} else {
			searchBy = append(searchBy, false)
		}
		totalHits, scrollID, results, err := search.PrototypeUserSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, scrollID)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		sendThis := make(map[string]interface{})
		sendThis["TotalHits"] = totalHits
		sendThis["ScrollID"] = scrollID
		sendThis["Results"] = results
		data, err := json.Marshal(sendThis)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		fmt.Fprintln(w, string(data))
	}
	if filter == `skills` {
		results, err := search.Skills(client.Eclient, strings.ToLower(query))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		data, err := json.Marshal(results)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		fmt.Fprintln(w, string(data))
	}
}
