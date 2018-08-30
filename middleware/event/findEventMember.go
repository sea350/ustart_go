package event

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sea350/ustart_go/globals"
	client "github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	term := r.FormValue("term")
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	if err == nil {
		var results []types.FloatingHead
		for _, element := range searchResults.Hits.Hits {
			head, err1 := uses.ConvertUserToFloatingHead(client.Eclient, element.Id)
			if err1 != nil {
				err = errors.New("there was one or more problems loading results")
				continue
			}
			results = append(results, head)
		}
		jsonnow, _ := json.Marshal(results)
		w.Write(jsonnow)
	}
}

//FindEventGuest ... find event guest
func FindEventGuest(w http.ResponseWriter, r *http.Request) {
	term := r.FormValue("term")
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	if err == nil {
		var results []types.FloatingHead
		for _, element := range searchResults.Hits.Hits {
			head, err1 := uses.ConvertUserToFloatingHead(client.Eclient, element.Id)
			if err1 != nil {
				err = errors.New("there was one or more problems loading results")
				continue
			}
			results = append(results, head)
		}
		jsonnow, _ := json.Marshal(results)
		w.Write(jsonnow)
	}
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	term := r.FormValue("term")
	var stringTerm []string
	stringTerm = strings.Split(term, ` `)
	query := elastic.NewBoolQuery()
	query = uses.MultiWildCardQuery(query, "Name", stringTerm, true)
	for _, element := range stringTerm {
		query = query.Should(elastic.NewFuzzyQuery("Name", strings.ToLower(element)).Fuzziness(2))
	}
	ctx := context.Background()
	searchResults, err := client.Eclient.Search().
		Index(globals.UserIndex).
		Query(query).
		Size(5).
		Do(ctx)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	if err == nil {
		var results []types.FloatingHead
		for _, element := range searchResults.Hits.Hits {
			head, err1 := uses.ConvertProjectToFloatingHead(client.Eclient, element.Id)
			if err1 != nil {
				err = errors.New("there was one or more problems loading results")
				continue
			}
			results = append(results, head)
		}
		jsonnow, _ := json.Marshal(results)
		w.Write(jsonnow)
	}
}
