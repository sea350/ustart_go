package main

import (
	"context"
	"io"
	"log"
	"strings"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//ScrollSuggestedUsers ...
//Scrolls through docs being loaded
var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func ScrollSuggestedUsers(eclient *elastic.Client, class int, tagArray []string, projects []types.ProjectInfo, followingUsers map[string]bool, userID string, scrollID string, majors []string, school string) (string, []types.FloatingHead, int, error) {

	ctx := context.Background()
	tags := make([]interface{}, 0)
	for tag := range tagArray {
		tags = append([]interface{}{strings.ToLower(tagArray[tag])}, tags...)
	}

	//Get mutual project members

	projectIDs := make([]interface{}, 0)
	for elements := range projects {
		projectIDs = append([]interface{}{strings.ToLower(projects[elements].ProjectID)}, projectIDs...)
	}

	followingUsers[userID] = true
	followIDs := make([]interface{}, 0)
	for id := range followingUsers {
		followIDs = append([]interface{}{id}, followIDs...)
	}

	majorsInterface := make([]interface{}, 0)
	for elements := range majors {
		majorsInterface = append([]interface{}{strings.ToLower(majors[elements])}, majorsInterface...)
	}

	suggestedUserQuery := elastic.NewBoolQuery()
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Tags", tags...)).Boost(2)
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Projects.ProjectID", projectIDs...)).Boost(1.5)
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Majors", majorsInterface...)).Boost(1.25)
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermQuery("UndergradSchool", school))
	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))
	suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Visible", true))
	suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Verified", true))
	suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Status", true))
	if class == 5 {
		suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermQuery("Class", 5))
	}

	//Please do not touch, very delicate
	var amt = 1

	if scrollID != `` {

		amt = 1
	} else {
		amt = 1

	}

	searchResults := eclient.Scroll().
		Index(globals.UserIndex).
		Query(suggestedUserQuery).
		Size(amt).
		Sort("_score", true)

	if len(scrollID) > 0 {
		searchResults = searchResults.ScrollId(scrollID)
	}

	res, err := searchResults.Do(ctx)

	if !(err == io.EOF && res != nil) && err != nil {
		if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}

		return "", nil, 0, err
	}

	// if res.Hits.TotalHits == 0 { //if no results just start recommending random
	// 	suggestedUserQuery = elastic.NewBoolQuery()
	// 	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Visible", true))
	// 	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Verified", true))
	// 	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Status", true))
	// 	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))
	// 	searchResults = eclient.Scroll().
	// 		Index(globals.UserIndex).
	// 		Query(suggestedUserQuery).
	// 		Size(1)

	// 	if len(scrollID) > 0 {
	// 		searchResults = searchResults.ScrollId(scrollID)
	// 	}

	// 	res, err = searchResults.Do(ctx)
	// 	if !(err == io.EOF && res != nil) && err != nil {
	// 		if err != io.EOF {
	// 			log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 			log.Println(err)
	// 		}
	// 		return "", nil, 0, err
	// 	}
	// }

	var heads []types.FloatingHead
	for _, hits := range res.Hits.Hits {
		newHead, err := uses.ConvertUserToFloatingHead(eclient, hits.Id)
		if err == nil {
			heads = append(heads, newHead)
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}

	}

	return res.ScrollId, heads, len(heads), err

}

// func sugg(eclient *elastic.Client, class int, tagArray []string, projects []types.ProjectInfo, followingUsers map[string]bool, userID string, scrollID string, majors []string, school string) (string, []types.FloatingHead, int, error) {

// 	ctx := context.Background()
// 	tags := make([]interface{}, 0)
// 	for tag := range tagArray {
// 		tags = append([]interface{}{strings.ToLower(tagArray[tag])}, tags...)
// 	}

// 	//Get mutual project members

// 	projectIDs := make([]interface{}, 0)
// 	for elements := range projects {
// 		projectIDs = append([]interface{}{strings.ToLower(projects[elements].ProjectID)}, projectIDs...)
// 	}

// 	followingUsers[userID] = true
// 	followIDs := make([]interface{}, 0)
// 	for id := range followingUsers {
// 		followIDs = append([]interface{}{id}, followIDs...)
// 	}

// 	majorsInterface := make([]interface{}, 0)
// 	for elements := range majors {
// 		majorsInterface = append([]interface{}{strings.ToLower(majors[elements])}, majorsInterface...)
// 	}

// 	suggestedUserQuery := elastic.NewBoolQuery()
// 	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))

// 	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Visible", true))
// 	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Verified", true))
// 	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Status", true))

// 	if class == 5 {
// 		suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermQuery("Class", 5))
// 	}

// 	if len(tagArray) == 0 && len(projects) == 0{
// 		maq := elastic.NewMatchAllQuery()
// 		suggestedUserQuery = suggestedUserQuery.Must(maq)

// 	}else if len(tagArray) == 0{
// 		projectQuery :=
// 	}

// 	suggestedUserQuery1 := elastic.NewBoolQuery()
// 	suggestedUserQuery1 = suggestedUserQuery1.Must(elastic.NewTermsQuery("Tags", tags...).Boost(1.75))
// 	suggestedUserQuery1 = suggestedUserQuery1.Must(elastic.NewTermsQuery("Projects.ProjectID", projectIDs...).Boost(2.5))

// 	// suggestedUserQuery2 := elastic.NewTermsQuery("Projects.ProjectID", projectIDs...).Boost(2.5)

// 	// suggestedUserQuery3 := elastic.NewTermsQuery("Majors", majorsInterface...).Boost(1.2)

// 	// suggestedUserQuery4 := elastic.NewTermQuery("UndergradSchool", school).Boost(2.0)

// 	suggestedUserQuery = suggestedUserQuery.Must(suggestedUserQuery1) //, suggestedUserQuery2) //, suggestedUserQuery3, suggestedUserQuery4)

// 	// suggestedUserQuery := elastic.NewBoolQuery()
// 	// suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Tags", tags...)).Boost(2)
// 	// suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Projects.ProjectID", projectIDs...)).Boost(1.5)
// 	// suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Majors", majorsInterface...)).Boost(1.25)
// 	// suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermQuery("UndergradSchool", school))
// 	// suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))
// 	// suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Visible", true))
// 	// suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Verified", true))
// 	// suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Status", true))
// 	// if class == 5 {
// 	// 	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermQuery("Class", 5))
// 	// }

// 	//Please do not touch, very delicate
// 	var amt = 1

// 	// if scrollID != `` {

// 	// 	amt = 1
// 	// } else {
// 	// 	amt = 1

// 	// }

// 	searchResults := eclient.Scroll().
// 		Index(globals.UserIndex).
// 		Query(suggestedUserQuery).
// 		Size(amt)

// 	if len(scrollID) > 0 {
// 		searchResults = searchResults.ScrollId(scrollID)
// 	}

// 	res, err := searchResults.Do(ctx)

// 	if !(err == io.EOF && res != nil) && err != nil {
// 		if err != io.EOF {
// 			log.SetFlags(log.LstdFlags | log.Lshortfile)
// 			log.Println(err)
// 		}

// 		return "", nil, 0, err
// 	}

// 	fmt.Println("Total Hits:", res.Hits.TotalHits)
// 	var heads []types.FloatingHead
// 	for _, hits := range res.Hits.Hits {
// 		newHead, err := uses.ConvertUserToFloatingHead(eclient, hits.Id)
// 		if err == nil {
// 			heads = append(heads, newHead)
// 		} else {
// 			log.SetFlags(log.LstdFlags | log.Lshortfile)
// 			log.Println(err)
// 			continue
// 		}

// 	}

// 	return res.ScrollId, heads, len(heads), err

// }

// func main() {
// 	unames := []string{"min", "ryanrozbiani", "nevets", "yh1112", "sx563", "julietr", "sc6094", "Carlos", "HeatherMT"}

// 	for _, uname := range unames {
// 		usrID, err := getUser.IDByUsername(eclient, uname)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		_, usrFoll, err := getFollow.ByID(eclient, usrID)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		usr, err := getUser.UserByUsername(eclient, uname)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		_, h1, _, err := sugg(eclient, usr.Class, usr.Tags, usr.Projects, usrFoll.UserFollowing, usrID, "", usr.Majors, usr.University)
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			for _, h := range h1 {
// 				fmt.Println(h.FirstName, h.LastName)
// 			}
// 		}
// 	}
// }
