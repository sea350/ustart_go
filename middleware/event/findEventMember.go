package event

import (
	"context"
	"encoding/json"
	"errors"

	"net/http"

	"strings"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/globals"
	client "github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
	elastic "github.com/olivere/elastic"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		return
	}
	p := bluemonday.UGCPolicy()
	term := p.Sanitize(r.FormValue("term"))
	var stringTerm []string
	stringTerm = strings.Split(term, ` `)
	query := elastic.NewBoolQuery()
	query = uses.MultiWildCardQuery(query, "Username", stringTerm, true)
	for _, element := range stringTerm {
		query = query.Should(elastic.NewFuzzyQuery("Username", strings.ToLower(element)).Fuzziness(2))
	}
	ctx := context.Background()
	searchResults, err := client.Eclient.Search().
		Index(globals.UserIndex).
		Query(query).
		Size(5).
		Do(ctx)

	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	if err == nil {
		var results []string
		for _, element := range searchResults.Hits.Hits {
			head, err1 := uses.ConvertUserToFloatingHead(client.Eclient, element.Id)
			if err1 != nil {
				err = errors.New("there was one or more problems loading results")
				continue
			}
			results = append(results, head.Username)
		}
		jsonnow, _ := json.Marshal(results)
		w.Write(jsonnow)
	}
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)

	}

	var results []string
	for _, element := range userstruct.Projects {
		head, err1 := uses.ConvertProjectToFloatingHead(client.Eclient, element.ProjectID)
		if err1 != nil {
			err = errors.New("there was one or more problems loading results")
			continue
		}
		results = append(results, head.FirstName)
	}
	jsonnow, _ := json.Marshal(results)
	w.Write(jsonnow)
}
