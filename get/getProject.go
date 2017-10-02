package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	"context"
	
	"encoding/json"
	
)
const PROJECT_INDEX="test-project_data"
const PROJECT_TYPE="PROJECT"


func GetProjectByID(eclient *elastic.Client, projectID string)(types.Project,error){
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT DOC ID)
	//RETURNS A types. Project AND AN error
	var proj types.Project
	
	ctx:=context.Background()

	searchResult, err := eclient.Get().
		Index(PROJECT_INDEX).
		Type(PROJECT_TYPE).
		Id(projectID).
		Do(ctx)
	if (err!=nil){return proj,err}
	

	
	err = json.Unmarshal(*searchResult.Source, &proj)
	//if (err!=nil){return proj,err}

	return proj, err
	
}

func GetProjectByURL(eclient *elastic.Client, projectURL string)(types.Project,error){
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT URL)
	//RETURNS A types.Project AND AN error
	ctx:=context.Background()
	termQuery := elastic.NewTermQuery("URL",projectURL)
	searchResult,err:=eclient.Search().
		Index(PROJECT_INDEX).
		Query(termQuery).
		Do(ctx)

	
	var result string
	var proj types.Project
	for _,element:=range searchResult.Hits.Hits{
	
		result = element.Id
		break
	}
	
	proj, _ = GetProjectByID(eclient,result)

	return proj, err
	
}



func GetProjectIDByURL(eclient *elastic.Client, projectURL string)(string,error){
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT URL)
	//RETURNS A types.Project AND AN error
	ctx:=context.Background()
	termQuery := elastic.NewTermQuery("URL",projectURL)
	searchResult,err:=eclient.Search().
		Index(PROJECT_INDEX).
		Query(termQuery).
		Do(ctx)

	
	var result string
	
	for _,element:=range searchResult.Hits.Hits{
	
		result = element.Id
		break
	}
	
	
	return result, err
	
}

func URLInUse(eclient *elastic.Client, projectURL string)(bool, error){
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT URL)
	//RETURNS A types.Project AND AN error
	ctx:=context.Background()
	termQuery := elastic.NewTermQuery("URL",projectURL)
	searchResult,err:=eclient.Search().
		Index(PROJECT_INDEX).
		Query(termQuery).
		Do(ctx)

	if (err!=nil){return true, err}
	
	
	if (searchResult.Hits.TotalHits > 0) {return true, nil}
	

	return false, nil
	
}

