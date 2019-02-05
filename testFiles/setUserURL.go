package main

import (
	"fmt"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

func changeAccURL() {
	fmt.Println("getting user by email")
	id, err := get.IDByUsername(client.Eclient, "username")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = post.UpdateUser(client.Eclient, id, "Username", "newusername")
	fmt.Println(err)
}

func main() {
	changeAccURL()
}
