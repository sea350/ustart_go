package search

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/search"
	"github.com/sea350/ustart_go/types"
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
		scrollID, results, err := search.PrototypeProjectSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, "")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		cs.ListOfHeads = results
		cs.ScrollID = scrollID
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
		/*
			if r.FormValue("searchbyprojectname") != `` {
				searchBy = append(searchBy, true)
			} else {
				searchBy = append(searchBy, false)
			}
		*/
		scrollID, results, err := search.PrototypeEventSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, "")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		cs.ListOfHeads = results
		cs.ScrollID = scrollID
	}
	if filter == `users` {
		fmt.Println(r.FormValue("searchbypersonname"))
		fmt.Println(r.FormValue("searchbyusername"))
		fmt.Println(r.FormValue("searchbyskills"))
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
		scrollID, results, err := search.PrototypeUserSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, "")
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		cs.ListOfHeads = results
		cs.ScrollID = scrollID
	}
	if filter == `skills` {
		results, err := search.Skills(client.Eclient, strings.ToLower(query))
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
		cs.ListOfHeads = results
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "search-nil", cs)
}
