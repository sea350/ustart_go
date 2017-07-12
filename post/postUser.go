package Post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    "github.com/sea350/ustart/types"
    "context"
	// "errors"
    "fmt"
)

const USER_INDEX = "test-user_data"
const USER_TYPE  = "USER"




func IndexUser(newAcc types.User){
    ctx := context.Background()
	eclient, err:= elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {fmt.Println(err)}
	exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
    if err != nil {
        //return err 
		fmt.Println(err)
    }
    if !exists {
        //return errors.New("Index does not exist")
		fmt.Println(err)
    }
    _, err = eclient.Index().
        Index(USER_INDEX).
        Type(USER_TYPE).
        BodyJson(newAcc).
        Do(ctx)
    if err != nil {
        //return err
		fmt.Println(err)
	 
    }
    //return nil
}

