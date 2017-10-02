package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	"github.com/sea350/ustart_go/types"
	get "github.com/sea350/ustart_go/get"
	"context"
	"errors"
	//"fmt"

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
                },
                "Username":{
                	"type":"keyword"
                },
                "AccCreation":{
                	"type": date"
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
		createIndex, Err := eclient.CreateIndex(USER_INDEX).BodyString(mapping).Do(ctx)

		if Err != nil {
			// Handle error
			_,_ = eclient.IndexExists(USER_INDEX).Do(ctx)
			panic(Err)
		}
		if !createIndex.Acknowledged {
		}


		return errors.New("Index does not exist")
	}

	_, Err := eclient.Index().
		Index(USER_INDEX).
		Type(USER_TYPE).
		BodyJson(newAcc).
		Do(ctx)

	if (Err!=nil){return Err}


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

func UpdateUser(eclient *elastic.Client, userID string, field string, newContent interface{})error {
	//CHANGES A SINGLE FIELD OF ES USER DOCUMENT(requires an elastic client pointer,
	//	the user DocID, the feild you wish to modify as a string,
	//	and what you want to change that field to as any type necessary)
	//RETURN AN error IF SUCESSFUL error = nil

	ctx := context.Background()

	exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}

	_, err = get.GetUserByID(eclient, userID)
	if (err!=nil){return err}

	_, err = eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(userID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	

	return err
}


func AppendToUser(eclient *elastic.Client, usrID string, field string, data interface{})error{return nil}//RETURN HERE
func RemoveFromUser(eclient *elastic.Client, usrID string, field string, idx int, data interface{})error{return nil}






func AppendCollReq(eclient *elastic.Client, usrID string, collegueID string, whichOne bool)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	if (whichOne == true){
		usr.SentCollReq = append(usr.SentCollReq,collegueID)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"SentCollReq": usr.SentCollReq}).
			Do(ctx)

		return CheckSentCollReq(eclient, usrID, collegueID, true, 0)
	}else{
		usr.ReceivedCollReq = append(usr.ReceivedCollReq,collegueID)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"ReceivedCollReq": usr.ReceivedCollReq}).
			Do(ctx)

		return CheckReceivedCollReq(eclient, usrID, collegueID, true, 0)

	}

}


func DeleteCollReq(eclient *elastic.Client, usrID string, colleagueID string, whichOne bool, idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	if (whichOne == true){
		usr.SentCollReq = append(usr.SentCollReq[:idx],usr.SentCollReq[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"SentCollReq": usr.SentCollReq}).
			Do(ctx)

		return CheckSentCollReq(eclient, usrID,colleagueID, false, idx)
	}else{
		usr.ReceivedCollReq = append(usr.ReceivedCollReq[:idx],usr.ReceivedCollReq[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"ReceivedCollReq": usr.ReceivedCollReq}).
			Do(ctx)

		return CheckReceivedCollReq(eclient, usrID, colleagueID, false, idx)

	}

}





func CheckSentCollReq(eclient *elastic.Client, usrID string, colleagueID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.SentCollReq{

				if (theDoc.SentCollReq[i]==colleagueID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendCollReq(eclient, usrID, colleagueID,true)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteCollReq(eclient, usrID, colleagueID, true, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}

func CheckReceivedCollReq(eclient *elastic.Client, usrID string, colleagueID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.ReceivedCollReq{

				if (theDoc.ReceivedCollReq[i]==colleagueID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendCollReq(eclient, usrID, colleagueID,false)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteCollReq(eclient, usrID, colleagueID, false, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}








func AppendColleague(eclient *elastic.Client, usrID string, colleagueID string)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	
	usr.Colleagues = append(usr.Colleagues,colleagueID)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"Colleagues": usr.Colleagues}).
		Do(ctx)

	return CheckColleagues(eclient, usrID, colleagueID, true, 0)
	
}


func DeleteColleague(eclient *elastic.Client, usrID string, colleagueID string,idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	
	usr.Colleagues = append(usr.Colleagues[:idx],usr.Colleagues[idx+1:]...)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"Colleagues": usr.Colleagues}).
		Do(ctx)
	
	return CheckColleagues(eclient, usrID,colleagueID, false,idx)
	

	
}





func CheckColleagues(eclient *elastic.Client, usrID string, colleagueID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.Colleagues{

				if (theDoc.Colleagues[i]==colleagueID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendColleague(eclient, usrID, colleagueID)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteColleague(eclient, usrID, colleagueID,idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}








func AppendMajorMinor(eclient *elastic.Client, usrID string, major_minor string, whichOne bool)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	if (whichOne == true){
	usr.Majors = append(usr.Majors,major_minor)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"Majors": usr.Majors}).
		Do(ctx)

	return CheckMajor(eclient, usrID, major_minor, true, 0)
	}else{
		usr.Minors = append(usr.Minors,major_minor)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"Minors": usr.Minors}).
			Do(ctx)

		return CheckMinor(eclient, usrID, major_minor, true, 0)

	}

}

func DeleteMajorMinor(eclient *elastic.Client, usrID string, major_minor string, whichOne bool, idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	if (whichOne == true){
		usr.Majors = append(usr.Majors[:idx],usr.Majors[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"Majors": usr.Majors}).
			Do(ctx)

		return CheckMajor(eclient, usrID, major_minor, false, idx)
	}else{
		usr.Minors = append(usr.Minors[:idx],usr.Minors[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"Minors": usr.Minors}).
			Do(ctx)

		return CheckMinor(eclient, usrID, major_minor, false, idx)

	}

}



func CheckMajor(eclient *elastic.Client, usrID string, major string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.Majors{

				if (theDoc.Majors[i]==major){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendMajorMinor(eclient, usrID, major,true)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteMajorMinor(eclient, usrID, major, true, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}


func CheckMinor(eclient *elastic.Client, usrID string, minor string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.Minors{

				if (theDoc.Minors[i]==minor){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendMajorMinor(eclient, usrID, minor, true)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteMajorMinor(eclient, usrID, minor, false, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}

func AppendFollow(eclient *elastic.Client, usrID string, followID string, whichOne bool)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	if (whichOne == true){
		usr.Following = append(usr.Following,followID)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"Following": usr.SentCollReq}).
			Do(ctx)

		return CheckFollowing(eclient, usrID, followID, true, 0)
	}else{
		usr.Followers = append(usr.Followers,followID)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"Followers": usr.Followers}).
			Do(ctx)

		return CheckFollowers(eclient, usrID, followID, true, 0)

	}

}


func DeleteFollow(eclient *elastic.Client, usrID string, followID string, whichOne bool, idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	if (whichOne == true){
		usr.Following = append(usr.Following[:idx],usr.Following[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"Following": usr.Following}).
			Do(ctx)

		return CheckFollowers(eclient, usrID,followID, false, idx)
	}else{
		usr.Followers = append(usr.Followers[:idx],usr.Followers[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"Followers": usr.Followers}).
			Do(ctx)

		return CheckFollowing(eclient, usrID, followID, false, idx)

	}

}





func CheckFollowers(eclient *elastic.Client, usrID string, followID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.Followers{

				if (theDoc.Followers[i]==followID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendFollow(eclient, usrID, followID,true)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteFollow(eclient, usrID, followID, true, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}



func CheckFollowing(eclient *elastic.Client, usrID string, followID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.Following{

				if (theDoc.Following[i]==followID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendFollow(eclient, usrID, followID,false)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteFollow(eclient, usrID, followID, false, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}




func AppendProjReq(eclient *elastic.Client, usrID string, projID string, whichOne bool)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	if (whichOne == true){
		usr.SentProjReq = append(usr.SentProjReq,projID)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"SentProjReq": usr.SentProjReq}).
			Do(ctx)

		return CheckSentCollReq(eclient, usrID, projID, true, 0)
	}else{
		usr.ReceivedProjReq = append(usr.ReceivedProjReq,projID)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"ReceivedProjReq": usr.ReceivedProjReq}).
			Do(ctx)

		return CheckReceivedCollReq(eclient, usrID, projID, true, 0)

	}

}


func DeleteProjReq(eclient *elastic.Client, usrID string, projID string, whichOne bool, idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	if (whichOne == true){
		usr.SentProjReq = append(usr.SentProjReq[:idx],usr.SentProjReq[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"SentProjReq": usr.SentProjReq}).
			Do(ctx)

		return CheckSentProjReq(eclient, usrID, projID, false, idx)
	}else{
		usr.ReceivedProjReq = append(usr.ReceivedProjReq[:idx],usr.ReceivedProjReq[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"ReceivedProjReq": usr.ReceivedProjReq}).
			Do(ctx)

		return CheckReceivedProjReq(eclient, usrID, projID, false, idx)

	}

}





func CheckSentProjReq(eclient *elastic.Client, usrID string, projID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.SentProjReq{

				if (theDoc.SentProjReq[i]==projID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendProjReq(eclient, usrID, projID,true)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteProjReq(eclient, usrID, projID, true, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}


func CheckReceivedProjReq(eclient *elastic.Client, usrID string, projID string, action bool, idx int)error{
	isAppended := false

	for isAppended == false{
			theDoc, err := get.GetUserByID(eclient,usrID)
			if (err!=nil) {return errors.New("User does not exist")}

			for i:=range theDoc.ReceivedProjReq{

				if (theDoc.ReceivedProjReq[i]==projID){
					if (action == true){
						isAppended = true
						return nil

						}else{
							isAppended = false
						}
				}
			}
		
			if (action == true && isAppended == false){
				checkErr := AppendProjReq(eclient, usrID, projID,false)
				if (checkErr != nil){return checkErr}

			}else if (action == false && isAppended == false){
				return nil
				}

		}

		if (action == false && isAppended == true){
			checkErr := DeleteProjReq(eclient, usrID, projID, false, idx)
			if (checkErr != nil){return checkErr}
		}

		return nil

}


func AppendLikedEntryID(eclient *elastic.Client, usrID string, entryID string)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	
	usr.LikedEntryIDs = append(usr.LikedEntryIDs,entryID)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"LikedEntryIDs": usr.LikedEntryIDs}).
		Do(ctx)

	return err
	
}


func DeleteLikedEntryID(eclient *elastic.Client, usrID string, entryID string,idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	
	usr.LikedEntryIDs = append(usr.LikedEntryIDs[:idx],usr.LikedEntryIDs[idx+1:]...)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"LikedEntryIDs": usr.LikedEntryIDs}).
		Do(ctx)
	
	return err
	

	
}



func AppendProject(eclient *elastic.Client, usrID string, proj types.ProjectInfo)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	
	usr.Projects = append(usr.Projects,proj)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"Projects": usr.Projects}).
		Do(ctx)

	return err
	
}

func AppendLink(eclient *elastic.Client, usrID string, link types.Link)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	
	usr.QuickLinks = append(usr.QuickLinks,link)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"QuickLinks": usr.QuickLinks}).
		Do(ctx)

	return err
	
}


func DeleteLink(eclient *elastic.Client, usrID string, link types.Link,idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	
	usr.QuickLinks = append(usr.QuickLinks[:idx],usr.QuickLinks[idx+1:]...)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"Quicklinks": usr.QuickLinks}).
		Do(ctx)
	
	return err
	

	
}




func AppendTag(eclient *elastic.Client, usrID string, tag string)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	
	usr.Tags = append(usr.Tags,tag)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"Tags": usr.Tags}).
		Do(ctx)

	return err
	
}


func DeleteTag(eclient *elastic.Client, usrID string, tag string,idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	
	usr.Tags = append(usr.Tags[:idx],usr.Tags[idx+1:]...)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"Tags": usr.Tags}).
		Do(ctx)
	
	return err
	

	
}




func AppendBlock(eclient *elastic.Client, usrID string, blockID string, whichOne bool)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	if (whichOne == true){
		usr.BlockedUsers = append(usr.BlockedUsers,blockID)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedUsers": usr.BlockedUsers}).
			Do(ctx)

		return err
		}else{
			usr.BlockedUsers = append(usr.BlockedUsers,blockID)

			_,err =  eclient.Update().
				Index(USER_INDEX).
				Type(USER_TYPE).
				Id(usrID).
				Doc(map[string]interface{}{"BlockedBy": usr.BlockedBy}).
				Do(ctx)

			return err
		}
	
}


func DeleteBlock(eclient *elastic.Client, usrID string, blockID string,idx int, whichOne bool)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	if (whichOne == true){
		usr.BlockedUsers = append(usr.BlockedUsers[:idx],usr.BlockedUsers[idx+1:]...)

		_,err =  eclient.Update().
			Index(USER_INDEX).
			Type(USER_TYPE).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedUsers": usr.BlockedUsers}).
			Do(ctx)
	
		return err
		}else{
			usr.BlockedBy = append(usr.BlockedBy[:idx],usr.BlockedBy[idx+1:]...)

			_,err =  eclient.Update().
				Index(USER_INDEX).
				Type(USER_TYPE).
				Id(usrID).
				Doc(map[string]interface{}{"BlockedBy": usr.BlockedBy}).
				Do(ctx)
	
			return err


		}
	

	
}


func AppendEntryID(eclient *elastic.Client, usrID string, entryID string)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	
	usr.EntryIDs = append(usr.EntryIDs,entryID)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"EntryIDs": usr.EntryIDs}).
		Do(ctx)

	return err
	
}


func DeleteEntryID(eclient *elastic.Client, usrID string, entryID string,idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	
	usr.EntryIDs = append(usr.EntryIDs[:idx],usr.EntryIDs[idx+1:]...)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"EntryIDs": usr.EntryIDs}).
		Do(ctx)
	
	return err
	

	
}

func AppendConvoID(eclient *elastic.Client, usrID string, convoID string)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if (err!=nil) {return errors.New("User does not exist")}

	
	usr.ConversationIDs = append(usr.ConversationIDs,convoID)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"ConversationIDs": usr.ConversationIDs}).
		Do(ctx)

	return err
	
}


func DeleteConvoID(eclient *elastic.Client, usrID string, convoID string,idx int)error{
	ctx:= context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if (err!=nil) {return errors.New("User does not exist")}
	
	
	usr.ConversationIDs = append(usr.ConversationIDs[:idx],usr.ConversationIDs[idx+1:]...)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"ConversationIDs": usr.ConversationIDs}).
		Do(ctx)
	
	return err
	

	
}





func AppendSearch(eclient *elastic.Client, usrID string, newSearch string) error{
	ctx:=context.Background()
	usr, err := get.GetUserByID(eclient,usrID)
	if(err!=nil){return err}

	usr.SearchHist = append(usr.SearchHist, newSearch)

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(usrID).
		Doc(map[string]interface{}{"SearchHist": usr.SearchHist}).
		Do(ctx)
	
	return err

	
}