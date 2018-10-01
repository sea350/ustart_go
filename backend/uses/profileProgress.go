package uses

import (
	getUser "github.com/sea350/ustart_go/backend/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
	//post "github.com/sea350/ustart_go/backend/post"
)

//ProgressCheck ... WIP DO NOT USE
/*
fucntion uses userByID to get the user data and check individual field against the score board
returns a percentage, to do list, and error
*/
func ProgressCheck(eclient *elastic.Client, userID string) (int, []string, error) {
	//query return userdata and error
	data, err := getUser.UserByID(eclient, userID)

	if err != nil {
		return 0, nil, err
	}

	//array to store to do items
	var toDo []string

	//score counter
	currScore := 0
	maxScore := 100

	if data.Gender != "" {
		currScore += 5
	} else {
		toDo = append(toDo, "Gender")
	}

	if data.Phone != "" {
		currScore += 5
	} else {
		toDo = append(toDo, "Phone")
	}

	/*
		STILL NEED TO FINISH THE SCORE SYSTEM
	*/

	//return percentage
	return currScore / maxScore, toDo, err
}
