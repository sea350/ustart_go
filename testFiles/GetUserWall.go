package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first user's username: ")
	username1, _ := reader.ReadString('\n')
	username1 = username1[:len(username1)-1]
	id1, err := get.IDByUsername(client.Eclient, username1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User docID: " + id1)

	fmt.Print("Enter viewer user's username: ")
	username2, _ := reader.ReadString('\n')
	username2 = username2[:len(username2)-1]
	id2, err := get.IDByUsername(client.Eclient, username2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User docID: " + id2)

	_, entries, total, err := scrollpkg.ScrollPageUser(client.Eclient, id1, id2, "")
	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Total entries: ", total)
	for _, entry := range entries {
		fmt.Println("//------------------")
		fmt.Println(entry.Element.Content)
	}
}
