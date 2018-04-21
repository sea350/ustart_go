package search

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/search"
	"github.com/sea350/ustart_go/uses"
)

//Page ... draws search page
func Page(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	cs := client.ClientSide{}

	var searchBy []bool

	query := r.FormValue("query")
	filter := r.FormValue("searchFilterGroup") //can be: all,users,projects

	//sortBy := r.FormValue("sortbyfilter")
	searchMajors := uses.ConvertStrToStrArr(r.FormValue("searchlistmajors"))
	fmt.Println(searchMajors)
	searchSkills := uses.ConvertStrToStrArr(r.FormValue("searchlistskills")) //array
	fmt.Println(searchSkills)
	//searchByLocationCountry := uses.ConvertStrToStrArr(r.FormValue("searchbylocationcountry"))
	//searchByLocationState := uses.ConvertStrToStrArr(r.FormValue("searchbylocationstate"))

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
		results, err := search.SearchProject(client.Eclient, strings.ToLower(query))
		if err != nil {
			fmt.Println("err: middleware/search/search line 26")
		}
		cs.ListOfHeads = results
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
		results, err := search.SearchProfile(client.Eclient, strings.ToLower(query))
		if err != nil {
			fmt.Println("err: middleware/search/search line 34")
		}
		cs.ListOfHeads = results
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "search-nil", cs)
}
