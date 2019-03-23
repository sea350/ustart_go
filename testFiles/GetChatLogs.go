package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
)

func isInt(input string) bool {
	if _, err := strconv.Atoi(input); err != nil {
		fmt.Println(input)
		return false
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter chat docID: ")
	id, _ := reader.ReadString('\n')
	id = id[:len(id)-1]

	var idx int
	fmt.Print("Enter index to start from: ")
	idxString, _ := reader.ReadString('\n')
	idxString = idxString[:len(idxString)-1]
	if idxString != "" {
		for !isInt(idxString) {
			fmt.Print("You did not enter a number, please try again: ")
			idxString, _ = reader.ReadString('\n')
			idxString = idxString[:len(idxString)-1]
			if idxString == "" {
				break
			}
		}
		number, err := strconv.Atoi(idxString)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}
		idx = number
	}

	var quant int
	fmt.Print("Enter number of messages: ")
	lenString, _ := reader.ReadString('\n')
	lenString = lenString[:len(lenString)-1]
	if lenString != "" {
		for !isInt(lenString) {
			fmt.Print("You did not enter a number, please try again: ")
			lenString, _ = reader.ReadString('\n')
			lenString = lenString[:len(lenString)-1]
			if lenString == "" {
				break
			}
		}
		number, err := strconv.Atoi(lenString)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			return
		}
		quant = number
	}

	_, msgs, err := uses.ChatLoad(client.Eclient, id, idx, quant)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	for _, msg := range msgs {
		fmt.Println(msg.SenderID)
		fmt.Println(msg.Content)
	}

}
