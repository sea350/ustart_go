package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	elastic "github.com/olivere/elastic"
	getFollow "github.com/sea350/ustart_go/get/follow"
	getUser "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//ScrollSuggestedUsers ...
//Scrolls through docs being loaded
var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func sugg(eclient *elastic.Client, class int, tagArray []string, projects []types.ProjectInfo, followingUsers map[string]bool, userID string, scrollID string, majors []string, school string) (string, []types.FloatingHead, int, error) {

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
	suggestedUserQuery0 := elastic.NewTermsQuery("Tags", tags...).Boost(5)
	// suggestedUserQuery1 := elastic.NewTermsQuery("Projects.ProjectID", projectIDs...).Boost(4)
	// suggestedUserQuery2 := elastic.NewTermsQuery("Majors", majorsInterface...).Boost(3)

	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermQuery("UndergradSchool", school))
	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))

	suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Visible", true))
	suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Verified", true))
	suggestedUserQuery = suggestedUserQuery.Filter(elastic.NewTermQuery("Status", true))

	if class == 5 {
		suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermQuery("Class", 5))
	}

	suggestedUserQuery = suggestedUserQuery.Must(suggestedUserQuery0) //, suggestedUserQuery1, suggestedUserQuery2)
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
		Size(amt)

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

func main() {

	// eclient *elastic.Client, class int, tagArray []string, projects []types.ProjectInfo,
	// followingUsers map[string]bool, userID string, scrollID string, majors []string, school string

	minID, err := getUser.IDByUsername(eclient, "min")
	if err != nil {
		fmt.Println(err)
	}

	ryanID, err := getUser.IDByUsername(eclient, "ryanrozbiani")
	if err != nil {
		fmt.Println(err)
	}
	stevenID, err := getUser.IDByUsername(eclient, "nevets")
	if err != nil {
		fmt.Println(err)
	}
	yunjieID, err := getUser.IDByUsername(eclient, "yh1112")
	if err != nil {
		fmt.Println(err)
	}
	/////////////////////////////////////////////////////////////////////////////////////////////////

	_, minFoll, err := getFollow.ByID(eclient, minID)
	if err != nil {
		fmt.Println(err)
	}
	_, ryanFoll, err := getFollow.ByID(eclient, ryanID)
	if err != nil {
		fmt.Println(err)
	}
	_, stevenFoll, err := getFollow.ByID(eclient, stevenID)
	if err != nil {
		fmt.Println(err)
	}
	_, yunjieFoll, err := getFollow.ByID(eclient, yunjieID)
	if err != nil {
		fmt.Println(err)
	}

	////////////////////////////////////////////////////////////////////////////////////////////////
	min, err := getUser.UserByUsername(eclient, "min")
	if err != nil {
		fmt.Println(err)
	}

	ryan, err := getUser.UserByUsername(eclient, "ryanrozbiani")
	if err != nil {
		fmt.Println(err)
	}
	steven, err := getUser.UserByUsername(eclient, "nevets")
	if err != nil {
		fmt.Println(err)
	}
	yunjie, err := getUser.UserByUsername(eclient, "yh1112")
	if err != nil {
		fmt.Println(err)
	}
	var minList = []string{}
	var ryanList = []string{}
	var stevenList = []string{}
	var yunjieList = []string{}

	s1a, h1, _, _ := sugg(eclient, min.Class, min.Tags, min.Projects, minFoll.UserFollowing, minID, "", min.Majors, min.University)
	s1b, h1b, _, _ := sugg(eclient, min.Class, min.Tags, min.Projects, minFoll.UserFollowing, minID, s1a, min.Majors, min.University)
	_, h1c, _, _ := sugg(eclient, min.Class, min.Tags, min.Projects, minFoll.UserFollowing, minID, s1b, min.Majors, min.University)
	if len(minList) > 0 {
		minList = append(minList, h1[0].FirstName, h1b[0].FirstName, h1c[0].FirstName)
		fmt.Println(minList)
	}

	s2a, h2, _, _ := sugg(eclient, ryan.Class, ryan.Tags, ryan.Projects, ryanFoll.UserFollowing, ryanID, "", ryan.Majors, ryan.University)
	s2b, h2b, _, _ := sugg(eclient, ryan.Class, ryan.Tags, ryan.Projects, ryanFoll.UserFollowing, ryanID, s2a, ryan.Majors, ryan.University)
	_, h2c, _, _ := sugg(eclient, ryan.Class, ryan.Tags, ryan.Projects, ryanFoll.UserFollowing, ryanID, s2b, ryan.Majors, ryan.University)
	if len(stevenList) > 0 {
		ryanList = append(ryanList, h2[0].FirstName, h2b[0].FirstName, h2c[0].FirstName)
		fmt.Println(ryanList)
	}

	s3a, h3, _, _ := sugg(eclient, steven.Class, steven.Tags, steven.Projects, stevenFoll.UserFollowing, stevenID, "", steven.Majors, steven.University)
	s3b, h3b, _, _ := sugg(eclient, steven.Class, steven.Tags, steven.Projects, stevenFoll.UserFollowing, stevenID, s3a, steven.Majors, steven.University)
	_, h3c, _, _ := sugg(eclient, steven.Class, steven.Tags, steven.Projects, stevenFoll.UserFollowing, stevenID, s3b, steven.Majors, steven.University)

	if len(stevenList) > 0 {
		stevenList = append(stevenList, h3[0].FirstName, h3b[0].FirstName, h3c[0].FirstName)
		fmt.Println(stevenList)
	}

	s4a, h4, _, _ := sugg(eclient, yunjie.Class, yunjie.Tags, yunjie.Projects, yunjieFoll.UserFollowing, yunjieID, "", yunjie.Majors, yunjie.University)
	s4b, h4b, _, _ := sugg(eclient, yunjie.Class, yunjie.Tags, yunjie.Projects, yunjieFoll.UserFollowing, yunjieID, s4a, yunjie.Majors, yunjie.University)
	_, h4c, _, _ := sugg(eclient, yunjie.Class, yunjie.Tags, yunjie.Projects, yunjieFoll.UserFollowing, yunjieID, s4b, yunjie.Majors, yunjie.University)

	if len(yunjieList) > 0 {
		yunjieList = append(yunjieList, h4[0].FirstName, h4b[0].FirstName, h4c[0].FirstName)
		fmt.Println(yunjieList)
	}

}

// eclient *elastic.Client, class int, tagArray []string, projects []types.ProjectInfo,
// followingUsers map[string]bool, userID string, scrollID string, majors []string, school string
