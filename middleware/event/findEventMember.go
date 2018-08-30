package event

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	jsonthis := []string{"Magic Johnson, Magic White, Magic Pomper"}
	jsonnow, _ := json.Marshal(jsonthis)
	fmt.Fprintln(w, jsonnow)
}

//FindEventGuest ... find event guest
func FindEventGuest(w http.ResponseWriter, r *http.Request) {
	jsonthis := []string{"Magic Johnson, Magic White, Magic Pomper"}
	jsonnow, _ := json.Marshal(jsonthis)
	fmt.Fprintln(w, jsonnow)
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	jsonthis := []string{"Magic Johnson, Magic White, Magic Pomper"}
	jsonnow, _ := json.Marshal(jsonthis)
	fmt.Fprintln(w, jsonnow)
}
