package main

import (
	"context"
	"fmt"
	"strings"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
	postProject "github.com/sea350/ustart_go/post/project"
	postUser "github.com/sea350/ustart_go/post/user"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

// newProj.Avatar = "https://i.imgur.com/8BnkFLO.png"
// newProj.Banner = "https://i.imgur.com/XTj1t1J.png"
// newUsr.Avatar = "https://i.imgur.com/8BnFLO.png"
// newUsr.Banner = "https://i.imgur.com/XTj1t1J.png"

// //////////////////////////////////////////////
// newProj.Avatar = "https://ustart-default.s3.amazonaws.com/Defult_Project_Page_Logo.png"
// newProj.Banner = "https://ustart-default.s3.amazonaws.com/Defult_Profile_Banner_Logo.png"
// newUsr.Avatar = "https://ustart-default.s3.amazonaws.com/Defult_Profile_Page_Logo.png"
// newUsr.Banner = "https://ustart-default.s3.amazonaws.com/Defult_Profile_Banner_Logo.png"

func main() {
	ctx := context.Background()

	queryUAvi := elastic.NewTermQuery("Avatar", strings.ToLower("https://i.imgur.com/8BnFLO.png"))

	res, err := eclient.Search().
		Index(globals.UserIndex).
		Query(queryUAvi).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
	for _, hit := range res.Hits.Hits {
		err = postUser.UpdateUser(eclient, hit.Id, "Avatar", "https://ustart-default.s3.amazonaws.com/Defult_Profile_Page_Logo.png")
		fmt.Println(hit.Id)
		if err != nil {
			fmt.Println("LINE 45,", err)
		}
	}
	//////////////////////////////////////////////////////////////////////////////////////////////
	queryPAvi := elastic.NewTermQuery("Avatar", strings.ToLower("https://i.imgur.com/8BnkFLO.png"))

	res, err = eclient.Search().
		Index(globals.ProjectIndex).
		Query(queryPAvi).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
	for _, hit := range res.Hits.Hits {
		err = postProject.UpdateProject(eclient, hit.Id, "Avatar", "https://ustart-default.s3.amazonaws.com/Defult_Project_Page_Logo.png")
		fmt.Println(hit.Id)
		if err != nil {
			fmt.Println("LINE 65,", err)
		}
	}
	//////////////////////////////////////////////////////////////////////////////
	queryUBan := elastic.NewTermQuery("Banner", strings.ToLower("https://i.imgur.com/XTj1t1J.png"))

	res, err = eclient.Search().
		Index(globals.UserIndex).
		Query(queryUBan).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
	for _, hit := range res.Hits.Hits {
		err = postUser.UpdateUser(eclient, hit.Id, "Banner", "https://ustart-default.s3.amazonaws.com/Defult_Profile_Banner_Logo.png")
		fmt.Println(hit.Id)
		if err != nil {
			fmt.Println("LINE 45,", err)
		}
	}
	/////////////////////////////////////////////////////////////
	queryPBan := elastic.NewTermQuery("Banner", strings.ToLower("https://i.imgur.com/XTj1t1J.png"))

	res, err = eclient.Search().
		Index(globals.ProjectIndex).
		Query(queryPBan).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}
	for _, hit := range res.Hits.Hits {
		err = postProject.UpdateProject(eclient, hit.Id, "Banner", "https://ustart-default.s3.amazonaws.com/Defult_Profile_Banner_Logo.png")
		fmt.Println(hit.Id)
		if err != nil {
			fmt.Println("LINE 65,", err)
		}
	}

}
