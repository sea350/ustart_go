package event

import (
	"fmt"
	"net/http"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Magic Johnson (Member)")
}

//FindEventGuest ... find event guest
func FindEventGuest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Magic Johnson (Guest)")
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Magic Johnson (Project)")
}
