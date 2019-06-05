package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	elastic "github.com/olivere/elastic"
	get "github.com/sea350/ustart_go/get/project"
	post "github.com/sea350/ustart_go/post/project"

	// getUser "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/globals"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

//Jv63yWgBN3Vvtvdiu5YP

func main() {
	//TestProject, ustartvipmarketing, hello

	urls := []string{"TestProject", "hello", "ustartvipmarketing"}
	for _, url := range urls {
		projID, err := get.ProjectIDByURL(eclient, "")
		if err != nil {
			fmt.Println(err)
		}
		err = post.InvisProject(eclient, projID)
		if err != nil {
			fmt.Println(err)
		}

	}

}
