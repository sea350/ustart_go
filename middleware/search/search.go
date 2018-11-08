package search

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

	fmt.Println("Filter: ", filter)
	fmt.Println("Query: ", query)
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
		numHits, scrollID, results, err := search.PrototypeProjectSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, "")
		if err != nil && err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		cs.ListOfHeads = results
		cs.ScrollID = scrollID
		cs.Hits = numHits
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
		numHits, scrollID, results, err := search.PrototypeEventSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, "")
		if err != nil && err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		cs.ListOfHeads = results
		cs.ScrollID = scrollID
		cs.Hits = numHits
	}
	if filter == `users` {
		fmt.Println("searchbypersonname: ", searchbypersonname)
		fmt.Println("searchbyusername: ", searchbyusername)
		fmt.Println("searchbyskills: ", searchbyskills)
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
		numHits, scrollID, results, err := search.PrototypeUserSearchScroll(client.Eclient, strings.ToLower(query), 0, searchBy, searchMajors, searchSkills, []types.LocStruct{}, "")
		if err != nil && err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		cs.ListOfHeads = results
		cs.ScrollID = scrollID
		cs.Hits = numHits
	}
	if filter == `skills` {
		results, err := search.Skills(client.Eclient, strings.ToLower(query))
		if err != nil && err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		cs.ListOfHeads = results
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "search-nil", cs)
}
