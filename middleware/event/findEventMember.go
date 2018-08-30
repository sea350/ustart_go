package event

import (
	"fmt"
	"net/http"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FindEventMember Start")
	fmt.Fprintln(w, "Magic Johnson (Member)")
	fmt.Println("FindEventMember End")
}

//FindEventGuest ... find event guest
func FindEventGuest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FindEventGuest Start")
	fmt.Fprintln(w, "Magic Johnson (Guest)")
	fmt.Println("FindEventGuest End")
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FindEventProject Start")
	fmt.Fprintln(w, "Magic Johnson (Project)")
	fmt.Println("FindEventProject End")
}
