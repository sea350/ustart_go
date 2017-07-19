package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart/types"
	"context"
	//"encoding/json"
	//"errors"
	"fmt"
	"reflect"
)

const USER_INDEX="test-user_data"
const USER_TYPE="USER"

func GetUserFromEmail(email string) (string, []types.User) {
	//SEARCHES ES FOR A CERTAIN USER (REQUIRES USER EMAIL STRING)
	//IF SUCESSFUL SHOULD RETURN USER ARRAY OF SIZE 1
	ctx := context.Background()

	eclient, anErr := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if anErr != nil {fmt.Println(anErr)}
	// userID := "TestMan"
	
	var ret []types.User
	exists, err:= eclient.IndexExists(USER_INDEX).Do(ctx) 
	
	if err != nil {
		//fmt.Println(err)
	}
	if !exists {
		//fmt.Println(err)
	}

	matchQuery := elastic.NewMatchQuery("Email",email)

	searchResult, err := eclient.Search().
		Index("test-user_data").
		Query(matchQuery).
		Do(ctx)

	// after here = good
	 if err != nil{fmt.Println(err)}
	 var ttyp types.User  // ttyp is a "t" of type User.
	 for _, item := range searchResult.Each(reflect.TypeOf(ttyp)){
	 	if t, ok := item.(types.User); ok {	
			fmt.Println(t)
 			ret = append(ret, t)
	 	}
	 }
	// fmt.Println(ret)
	 if (ret.SIZE != 1) {

	 	err := "More than one user found"

	 	fmt.Println(err)
	 }

	 return err, ret
}