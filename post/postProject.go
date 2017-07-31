package post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    "github.com/sea350/ustart_go/types"
    "context"
	//"errors"
    "fmt"
)

const USER_INDEX = "test-project_data"
const USER_TYPE  = "PROJECT"

func IndexUser(eclient *elastic.Client, newAcc types.User)error {
	//ADDS NEW USER TO ES RECORDS (requires a User type)
    ctx := context.Background()
	
	exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
    if err != nil {
        return err 
		
    }
    if !exists {
        return errors.New("Index does not exist")
		//fmt.Println(err)
    }
    _, err = eclient.Index().
        Index(USER_INDEX).
        Type(USER_TYPE).
        BodyJson(newAcc).
        Do(ctx)
    if err != nil {
        return err
		
	 
    }
    return nil
}

func UpdateProject(eclient *elastic.Client, projectID string, projectPage types.User)error {
    //ADDS NEW USER TO ES RECORDS (requires a User type)
    ctx := context.Background()
    
    
    exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
    if err != nil {
        return err 
        //fmt.Println(err)
    }
    if !exists {
        return errors.New("Index does not exist")
        //fmt.Println(err)
    }
    _, err = eclient.Index().
        Index(USER_INDEX).
        Type(USER_TYPE).
        Id(projectID).
        BodyJson(projectPage).
        Do(ctx)
    if err != nil {
        return err
        
     
    }
    return nil
}


func ModifyDescription(eclient *elastic.Client, projectID string, newDescription string)error{
   
    ctx:=context.Background()

    proj, err:= eclient.Get().
        Index(USER_INDEX).
        Type(USER_TYPE).
        Id(projectID).
        Do(ctx)
    if (err != nil){
        return err
    }

    proj.Description = newDescription
    
    return UpdateUser(eclient,projectID,proj)

   
}








