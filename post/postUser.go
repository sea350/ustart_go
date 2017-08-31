package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	"github.com/sea350/ustart_go/types"
	get "github.com/sea350/ustart_go/get"
	"context"
	"errors"
	"fmt"

)

const USER_INDEX = "test-user_data"
const USER_TYPE  = "USER"

const mapping = `
{
    "mappings":{
        "User":{
            "properties":{
                "Email":{
                    "type":"keyword"
                }
                
            }
        }
    }
}`

func IndexUser(eclient *elastic.Client, newAcc types.User)error {
	//ADDS NEW USER TO ES RECORDS (requires an elastic client and a User type)
	//RETURNS AN error IF SUCESSFUL error = nil
	ctx := context.Background()

	exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
	if err != nil {return err}

	if !exists {
		fmt.Println("LINE 115")
		createIndex, Err := eclient.CreateIndex(USER_INDEX).BodyString(mapping).Do(ctx)

		if Err != nil {
			// Handle error
			fmt.Println("LINE 120 DOES IT EXIST?")
			nowExists,_ := eclient.IndexExists(USER_INDEX).Do(ctx)
			fmt.Println(nowExists)
			panic(Err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("LINE 80")
		}


		return errors.New("Index does not exist")
	}

	_, Err := eclient.Index().
		Index(USER_INDEX).
		Type(USER_TYPE).
		BodyJson(newAcc).
		Do(ctx)

	if (Err!=nil){}


	return nil
}

func ReindexUser(eclient *elastic.Client, userID string, userAcc types.User)error {
	//ADDS NEW USER TO ES RECORDS (requires an elastic client pointer and a User type)
	//RETURN AN error IF SUCESSFUL error = nil

	ctx := context.Background()

	exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
	if err != nil {return err}

	if !exists {return errors.New("Index does not exist")}

	_, err = eclient.Index().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(userID).
		BodyJson(userAcc).
		Do(ctx)

	if err != nil {return err}

	return nil
}

func UpdateUser(eclient *elastic.Client, userID string,  newContent interface{}, field string)error {
	//ADDS NEW USER TO ES RECORDS (requires an elastic client pointer and a User type)
	//RETURN AN error IF SUCESSFUL error = nil

	ctx := context.Background()

	exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}

	_, err = get.GetUserById(eclient, userID)
	if (err!=nil){return err}

	_, err = eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(userID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	//if err != nil {return err}

	return nil
}

