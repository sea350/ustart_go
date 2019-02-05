package main

import (
	"fmt"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
)

func changeProjURL() {
	fmt.Println("getting user by email")
	id, err := get.ProjectIDByURL(client.Eclient, "username")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = post.UpdateProject(client.Eclient, id, "URLName", "newusername")
	fmt.Println(err)
}

func main() {
	changeAccURL()
}
