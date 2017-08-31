package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	"errors"
	"context"
	"encoding/json"
)

const USER_INDEX="test-user_data"
const USER_TYPE="USER"



func GetUserById(eclient *elastic.Client, userID string)(types.User, error){
	ctx:=context.Background() //intialize context background
	var usr types.User //initialize type user
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(userID).
		Do(ctx)

	if (err!=nil){return usr, err}

	Err:= json.Unmarshal(*searchResult.Source, &usr) //unmarshal type RawMessage into user struct

	return usr, Err

}


func EmailToUsername(email string)(string){
    
    var usr []rune //in case of special characters, use runes
    
    for _,element := range email{ //iterate through email string
        if(element!='@'){ //remove '@'
            usr=append(usr,element)
        } else {
        	usr=append(usr,'.') //replace '@' with '.'
        }
     }

    retUsr:=string(usr) //converts to string for username

    return retUsr //return s username

}




func GetUserByEmail(eclient *elastic.Client, email string)(types.User,error){

	ctx:=context.Background()

	//username:= EmailToUsername(email) //for username query
	termQuery := elastic.NewTermQuery("Email",email)
	searchResult,err:=eclient.Search().
		Index(USER_INDEX).
		Query(termQuery).
		Do(ctx)

	
	var result string
	var usr types.User
	for _,element:=range searchResult.Hits.Hits{
	
		result = element.Id
		break
	}
	
	usr, _ = GetUserById(eclient,result)

	return usr, err

}


func GetUserIDByEmail(eclient *elastic.Client, email string)(string,error){

	ctx:=context.Background()

	//username:= EmailToUsername(email)

	//termQuery := elastic.NewTermQuery("Username",username)

	termQuery := elastic.NewTermQuery("Email",email)
	searchResult,err:=eclient.Search().
		Index(USER_INDEX).
		Query(termQuery).
		Do(ctx)

		if(err!=nil){panic(err)}


	var result string //save id to a result variable
	if (searchResult.TotalHits() > 1) {return result,errors.New("More than one user found")} 
	
	for _,element:=range searchResult.Hits.Hits{ //interate through hits, get the element id
		result = element.Id
		
	}

		
	return result, err //return id

}

func EmailInUse(eclient *elastic.Client, theEmail string)(bool,error){
	ctx := context.Background()
	//username:=EmailToUsername(theEmail)
	//termQuery := elastic.NewTermQuery("Username",username)
	termQuery := elastic.NewTermQuery("Email", theEmail)
	searchResult,err:=eclient.Search().
		Index(USER_INDEX).
		Query(termQuery).
		Do(ctx)



	if(err!=nil){return false,err} //email might not be in use, but it's still an error

	exists:=searchResult.TotalHits() > 0 
	
	
	return exists, err


}

