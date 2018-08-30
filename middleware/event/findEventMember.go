package event

import (
	"fmt"
	"net/http"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Magic Johnson")
}

//FindEventGuest ... find event guest
func FindEventGuest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Magic Johnson")
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Magic Johnson")
}
