package search

import(

	elastic "gopkg.in/olivere/elastic.v5"
	//types "github.com/sea350/ustart_go/types"
	"context"
	//get "github.com/sea350/ustart_go/get"
	//"encoding/json"
	//"errors"
	"fmt"
	"strings"

)


var eclient, _= elastic.NewClient(elastic.SetURL("http://localhost:9200"))


const USER_INDEX = "test-user_data"
const USER_TYPE = "USER"

const PROJECT_INDEX = "test-project_data"
const PROJECT_TYPE  = "PROJECT"

const ENTRY_INDEX = "test-entry_data"
const ENTRY_TYPE = "ENTRY"




//only a mockup of code, certain things like iterating through searchResult.Hits aren't coded coded yet
func USearch(eclient *elastic.Client, searchFor []interface{}, queryType string, filter string) ([]interface{}, error){
	ctx:= context.Background()


	var results []interface{}
	if (strings.ToLower(filter) != "general"){return USearchByFilter(eclient,searchFor,queryType,filter)}

	if (strings.ToLower(queryType) == "match"){
		for s := range searchFor {

			query := elastic.NewMatchQuery("_all", searchFor[s])

			searchResult, err := eclient.Search().
				Query(query).
				Do(ctx)

			fmt.Println("Before Check")
			if (err !=nil){return results, err}
			if (searchResult.Hits.Hits != nil){
				fmt.Println("Something Appended")
				for _,element := range searchResult.Hits.Hits{
					results = append(results,element.Id)

					fmt.Println("Print ID")
					fmt.Println(element.Id)
				}
			}
		}

	}else{
		fmt.Println("LINE 63")
		for s:= range searchFor{
			query := elastic.NewTermQuery("_all", searchFor[s])

			searchResult, err := eclient.Search().
				Query(query).
				Do(ctx)


			fmt.Println("Before Check")
			if (err !=nil){return results, err}
			if (searchResult.Hits != nil){
				fmt.Println("Something Appended")
				for _,element := range searchResult.Hits.Hits{
					results = append(results,element.Id)

					fmt.Println("Print ID")
					fmt.Println(element.Id)
				}
			}
		}
	}

	return results, nil


}

func USearchByFilter(eclient *elastic.Client, searchFor []interface{}, queryType string, filter string) ([]interface{}, error){
	return searchFor, nil

}