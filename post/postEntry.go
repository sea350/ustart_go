package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	get "github.com/sea350/ustart_go/get"
	"context"
	"errors"
	"time"
)

const ENTRY_INDEX="test-entry_data"
const ENTRY_TYPE="ENTRY"

func IndexEntry(eclient *elastic.Client, newEntry types.Entry)(string, error) {
	//ADDS NEW ENTRY TO ES RECORDS (requires an elastic client and a User type)
	//RETURNS AN error IF SUCESSFUL error = nil
	ctx := context.Background()
	var entryID string

	idx, Err := eclient.Index().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		BodyJson(newEntry).
		Do(ctx)

	
	if (Err!=nil){return entryID,Err}
	entryID = idx.Id

	return entryID, nil
}




func ReindexEntry(eclient *elastic.Client, oldEntry types.Entry, entryID string)error{
	ctx:=context.Background()

	exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)

	if err != nil {return err}

	if !exists {return errors.New("Index does not exist")}

	_, err = eclient.Index().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		BodyJson(oldEntry).
		Do(ctx)

	if err != nil {return err}

	return nil
}

func UpdateEntryContent(eclient *elastic.Client, entryID string, newContent []rune) error{
	ctx:=context.Background()
	//stringified := string(newContent)

	exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}
	




	//script := elastic.NewScript("ctx._source.Content = newCont").Params(map[string]interface{}{"newCont": "I AM NEW"})
	message := newContent
	_, err = eclient.Update().Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"Content": message}).
		Do(ctx)


	
	if err != nil {return err}
	return nil
}


func UpdateEntry(eclient *elastic.Client, entryID string, field string, newContent interface{}) error{
	ctx:=context.Background()
	//stringified := string(newContent)


	exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}

	_, err=get.GetEntryByID(eclient, entryID)
	if (err!=nil){return err}

	
	//script := elastic.NewScript("ctx._source.Content = newCont").Params(map[string]interface{}{"newCont": "I AM NEW"})
	

	_, err = eclient.Update().Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)


	
	if err != nil {return err}
	return nil
}


func AppendLike(eclient *elastic.Client, entryID string, likerID string)(error){
	ctx:=context.Background()
	anEntry, err := get.GetEntryByID(eclient,entryID)
	if (err!=nil){return nil}
	newLike:=types.Like{}
	newLike.UserID = likerID
	newLike.TimeStamp = time.Now()
	anEntry.Likes = append(anEntry.Likes, newLike)
	_,err = eclient.Update().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"Likes": anEntry.Likes}).
		Do(ctx)



	return CheckLike(eclient, entryID,newLike,true)

}


func CheckLike(eclient *elastic.Client, entryID string, theLike types.Like, action bool) error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetEntryByID(eclient,entryID)
			if (err!=nil) {return errors.New("Entry does not exist")}

			for i:=range theDoc.Likes{

				if (theDoc.Likes[i]==theLike){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendLike(eclient, entryID, theLike.UserID)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteLike(eclient, entryID, theLike.UserID)
			if (checkErr != nil){return checkErr}
		}

		return nil
}


func DeleteLike(eclient *elastic.Client, entryID string, likerID string)error{
	ctx:=context.Background()
	anEntry, err := get.GetEntryByID(eclient,entryID)
	


	idx:=0
	for i:= range anEntry.Likes{
		if (likerID == anEntry.Likes[i].UserID){
			idx = i
		}
	}

	
	anEntry.Likes = append(anEntry.Likes[:idx],anEntry.Likes[idx+1:]...)

	_,err =  eclient.Update().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"Likes": anEntry.Likes}).
		Do(ctx)

	if(err!=nil){return err}
	var like types.Like
	like.UserID = likerID
	return CheckLike(eclient,entryID,like,false)


	
}


func AppendShareID(eclient *elastic.Client, entryID string, shareID string)error{
	ctx:= context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)


	anEntry.ShareIDs = append(anEntry.ShareIDs,shareID)

	_,err =  eclient.Update().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": anEntry.ShareIDs}).
		Do(ctx)

	if(err!=nil){return err}

	return CheckShareID(eclient,entryID,shareID,true,0)

}



func CheckShareID(eclient *elastic.Client, entryID string, shareID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetEntryByID(eclient,entryID)
			if (err!=nil) {return errors.New("Entry does not exist")}

			for i:=range theDoc.ShareIDs{

				if (theDoc.ShareIDs[i]==shareID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendShareID(eclient, entryID, shareID)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteShareID(eclient, entryID, shareID, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}


func DeleteShareID(eclient *elastic.Client, entryID string, shareID string, idx int)error{
	ctx:= context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)
	if (err!=nil){return nil}
	

	for i:= range anEntry.ShareIDs{
		if (shareID == anEntry.ShareIDs[i]){
			idx = i
		}
	}

	
	anEntry.ShareIDs = append(anEntry.ShareIDs[:idx],anEntry.ShareIDs[idx+1:]...)

	_,err =  eclient.Update().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": anEntry.ShareIDs}).
		Do(ctx)

	return CheckShareID(eclient,entryID,shareID,false,idx)

}



func DeleteReplyID(eclient *elastic.Client, entryID string, replyID string,idx int )error{
	ctx:= context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)

	for i:= range anEntry.ShareIDs{
		if (replyID == anEntry.ReplyIDs[i]){
			idx = i
		}
	}

	anEntry.ReplyIDs = append(anEntry.ReplyIDs[:idx],anEntry.ReplyIDs[idx+1:]...)

	_,err =  eclient.Update().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs":anEntry.ReplyIDs}).
		Do(ctx)

	if(err!=nil){return err}

	return CheckReplyID(eclient,entryID,replyID,false,idx)

}


func AppendReplyID(eclient *elastic.Client, entryID string, replyID string)error{
	//newLike.TimeStamp = time.Now()
	ctx:= context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)
	anEntry.ReplyIDs = append(anEntry.ReplyIDs, replyID)
	_,err = eclient.Update().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs": anEntry.ReplyIDs}).
		Do(ctx)

	if(err!=nil){return err}

	return CheckReplyID(eclient,entryID,replyID,true,0)


}



func CheckReplyID(eclient *elastic.Client, entryID string, replyID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetEntryByID(eclient,entryID)
			if (err!=nil) {return errors.New("Entry does not exist")}

			for i:=range theDoc.ReplyIDs{

				if (theDoc.ReplyIDs[i]==replyID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendReplyID(eclient, entryID, replyID)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteReplyID(eclient, entryID, replyID,idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}



















/*func AppendToEntry(eclient *elastic.Client, entryID string, newContent interface{}, field string) error{
	 ctx:=context.Background()

    exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)
    if err != nil {return err}
    if !exists {return errors.New("Index does not exist")}

    //script := elastic.NewScript("ctx._source."+field+"+= newCont").Params(map[string]interface{}{"newCont": newContent})
    script := elastic.NewScript("ctx._source."+field+".add(Params."+field+")").Params(map[string]interface{}{"newCont": newContent})
    _, err = eclient.Update().
    	Index(ENTRY_INDEX).
    	Type(ENTRY_TYPE).
    	Id(entryID).
        Script(script).
        Do(ctx)
	

	_, err = eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		ID(userID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	
    if err != nil {return err}

    return nil
}*/
/*
func AppendToEntryArrayField(eclient *elastic.Client, entryID string, newContent interface{}, field string) error{

	ctx:=context.Background()



	if (field == "Likes"){
		newCont:= types.Like(newContent)
		

		theEntry,err:=get.GetEntryByID(eclient,entryID)
		theEntry.Likes=append(theEntry.Likes,newCont)
		if (err!=nil){errors.New("Entry does not exist")}

		res,err := eclient.Update().
			Index(ENTRY_INDEX).
			Type(ENTRY_TYPE).
			Id(entryID).
			Doc(map[string]interface{}{"Likes": testEntry.Likes}).
			Do(ctx)
	
		if (err!=nil){return err}

		isUpdated:=false

		for isUpdated == false{
		
			theDoc,_ := get.GetEntryByID(eclient,entryID)
			for _,element:=range theDoc.Likes{
				if element==newContent{
					isUpdated = true
				}else{
					isUpdated=false
				} 
			}
			_, _ = elastic.NewUpdateService(eclient).
    			Index(ENTRY_INDEX).
    			Type(ENTRY_TYPE).
    			Id(entryID).
        		Doc(map[string]interface{}{"Likes": testEntry.Likes}).
        		Do(ctx)
        
		}
	}else if (field == ShareIDs){
		testEntry,err := get.GetEntryByID(eclient,entryID)
		testEntry.ShareIDs = append(testEntry.ShareIDs,string(newContent))
		res,_ := eclient.Update().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": testEntry.ShareIDs}).
		Do(ctx)
	

	isUpdated:=false

	for isUpdated == false{
		theDoc,_ := get.GetEntryByID(eclient,entryID)
		for _,element:=range theDoc.Likes{
			if element==newContent{
				isUpdated = true
			}else{
				isUpdated=false
			} 
		}
		_, _ = elastic.NewUpdateService(eclient).
    		Index(ENTRY_INDEX).
    		Type(ENTRY_TYPE).
    		Id(entryID).
        	Doc(map[string]interface{}{"ShareIDs": testEntry.ShareIDs}).
        	Do(ctx)
        
	}
	}else if (field == ReplyIDs){
		testEntry,err:=get.GetEntryByID(eclient,entryID)
		testEntry.ReplyIDs:=append(testEntry.ReplyIDs,string(newContent))
		res,_ := eclient.Update().
		Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs": testEntry.ReplyIDs}).
		Do(ctx)
	

	isUpdated:=false

	for isUpdated == false{
		theDoc,_ := get.GetEntryByID(eclient,entryID)
		for _,element:=range theDoc.Likes{
			if element==newContent{
				isUpdated = true
			}else{
				isUpdated=false
			} 
		}
		_, _ = elastic.NewUpdateService(eclient).
    		Index(ENTRY_INDEX).
    		Type(ENTRY_TYPE).
    		Id(entryID).
        	Doc(map[string]interface{}{"ReplyIDs": testEntry.ReplyIDs}).
        	Do(ctx)
        

	}
}

*/





/*
func RemoveFromEntry(eclient *elastic.Client, entryID string, field interface{})(error){
	ctx:=context.Background()
	//stringified := string(newContent)


	exists, err := eclient.IndexExists(ENTRY_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}
	
	var theEntry types.Entry

	theEntry, err=get.GetEntryByID(eclient, entryID)
	 
	if (err!=nil){return err}

	cmd:="ctx._source."+field+".remove("+field+")"
	key:=theEntry.field
	script:= elastic.NewScript(cmd).Params(map[string]interface{}{field:key}) 
	//message := "Hello"

	_, err = eclient.Update().Index(ENTRY_INDEX).
		Type(ENTRY_TYPE).
		ID(entryID).
		Script(script).
		Do(ctx)


	
	if err != nil {return err}
	return nil
}	
*/

/*func AddLike(eclient *elastic.Client, entryID string, newLike types.Like)error{

	update, err := client.Update().Index(ENTRY_INDEX).Type(ENTRY_TYPE).ID(entryID).
		Script("ctx._source.Likes += num").
		ScriptParams(map[string]interface{}{"num": 1}).
		Upsert(map[string]interface{}{"retweets": 0}).
		Do()

	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("New version of tweet %q is now %d", update.ID, update.Version)
}*/