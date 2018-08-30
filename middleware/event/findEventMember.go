package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sea350/ustart_go/globals"
	client "github.com/sea350/ustart_go/middleware/client"
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
		Pretty(true).
		Do(ctx)

	fmt.Println("Debugging MemberFind for " + term + ": " + string(searchResults.Hits.TotalHits) + " results found")

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	if err == nil {
		jsonnow, _ := json.Marshal(searchResults.Hits.Hits)
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
		Pretty(true).
		Do(ctx)

	fmt.Println("Debugging MemberFind for " + term + ": " + string(searchResults.Hits.TotalHits) + " results found")

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	if err == nil {
		jsonnow, _ := json.Marshal(searchResults.Hits.Hits)
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
		Pretty(true).
		Do(ctx)

	fmt.Println("Debugging MemberFind for " + term + ": " + string(searchResults.Hits.TotalHits) + " results found")

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	if err == nil {
		jsonnow, _ := json.Marshal(searchResults.Hits.Hits)
		w.Write(jsonnow)
	}
}
