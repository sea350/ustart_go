package post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    "github.com/sea350/ustart_go/types"
    "context"
    "errors"
)

const USER_INDEX = "test-user_data"
const USER_TYPE  = "USER"

func IndexUser(eclient *elastic.Client, newAcc types.User)error {
	//ADDS NEW USER TO ES RECORDS (requires an elastic client and a User type)
    //RETURNS AN error IF SUCESSFUL error = nil
    ctx := context.Background()
	
    exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)

    if err != nil {return err}

    if !exists {return errors.New("Index does not exist")}

    _, err = eclient.Index().
        Index(USER_INDEX).
        Type(USER_TYPE).
        BodyJson(newAcc).
        Do(ctx)

    if err != nil {fmt.Println(err)}

    return nil
}

func UpdateUser(eclient *elastic.Client, userID string, userAcc types.User)error {
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


func ModifyDescription(eclient *elastic.Client, userID string, newDescription string){
   
    ctx:=context.Background()

    usr:= eclient.Get().
        Index(USER_INDEX).
        Type(USER_TYPE).
        Id(userID).
       // BodyJson(userAcc).
        Do(ctx)

    usr.Description = newDescription
    
    UpdateUser(eclient,userID,usr)
}








