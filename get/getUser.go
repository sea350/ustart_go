package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	"context"
	"reflect"
)

const USER_INDEX="test-user_data"
const USER_TYPE="USER"

func GetUserFromEmail(eclient *elastic.Client, email string) ([]types.User, error) {
	//SEARCHES ES FOR A CERTAIN USER (REQUIRES USER EMAIL STRING)
	//IF SUCCESSFUL SHOULD RETURN []types.User OF SIZE 1 AND AN error
	ctx := context.Background()
	var ret []types.User

	matchQuery := elastic.NewMatchQuery("Email",email)
	
	searchResult, err := eclient.Search().
		Index(USER_INDEX).
		Query(matchQuery).
		Do(ctx)

	if err != nil{return ret, err}

	var ttyp types.User  // ttyp is a "t" of type User.
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)){
		if t, ok := item.(types.User); ok {	
			ret = append(ret, t)
		}
	}

	if (err != nil) {return ret, err}
	return ret, err
}

func GetIdFromEmail(eclient *elastic.Client, email string) ([]string, error) {
	//SEARCHES ES FOR A CERTAIN USER (REQUIRES elastic client pointer AND A string USER EMAIL)
	//IF SUCCESSFUL SHOULD RETURN string array OF SIZE 1 AND error
	ctx := context.Background()
	var ids []string
	
	exists, err:= eclient.IndexExists(USER_INDEX).Do(ctx) 
	
	if err != nil {return ids, err}
	if !exists {return ids, err}

	matchQuery := elastic.NewMatchQuery("Email", email)

	searchResult, err := eclient.Search().
		Index(USER_INDEX).
		Query(matchQuery).
		Do(ctx)


	if err != nil{return ids, err}

	for _, hit := range searchResult.Hits.Hits{
		ids = append(ids, hit.Id)
	}

	return ids, err
}

func GetUserFromId(eclient *elastic.Client, userID string)(types.User, error){
	//SEARCHES ES FOR A CERTAIN USER (REQUIRES elastic client pointer AND A 
	//		string USER EMAIL)
	//IF SUCCESSFUL SHOULD RETURN string array OF SIZE 1 AND error
	ctx:=context.Background()
	searchResult, err := eclient.Get().
		Index(USER_INDEX).
        Type(USER_TYPE).
        Id(userID).
        Do(ctx)
	
	//if (err!=nil) {return , err}


    usr:=searchResult

	return usr, err
}