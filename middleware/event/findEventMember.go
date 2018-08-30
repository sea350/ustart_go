package event

import (
	"fmt"
	"net/http"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Magic Johnson (Member)")
	fmt.Fprintln(w, "Magic King (Member)")
	fmt.Fprintln(w, "Magic Queen (Member)")
	fmt.Fprintln(w, "Magic Rook (Member)")
}

//FindEventGuest ... find event guest
func FindEventGuest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Magic Johnson (Guest)")
	fmt.Fprintln(w, "Magic Pawn (Guest)")
	fmt.Fprintln(w, "Magic Octopus (Guest)")
	fmt.Fprintln(w, "Magic Histu (Guest)")
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Magic Johnson (Project)")
	fmt.Fprintln(w, "Magic Michael (Project)")
	fmt.Fprintln(w, "Magic Pencil (Project)")
	fmt.Fprintln(w, "Magic Jason (Project)")
}
