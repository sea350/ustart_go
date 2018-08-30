package event

import (
	"encoding/json"
	"net/http"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	jsonthis := []string{"Magic Johnson", "Magic White", "Magic Pomper"}
	jsonnow, _ := json.Marshal(jsonthis)
	w.Write(jsonnow)
}

//FindEventGuest ... find event guest
func FindEventGuest(w http.ResponseWriter, r *http.Request) {
	jsonthis := []string{"Magic Johnson", "Magic White", "Magic Pomper"}
	jsonnow, _ := json.Marshal(jsonthis)
	w.Write(jsonnow)
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	jsonthis := []string{"Magic Johnson", "Magic White", "Magic Pomper"}
	jsonnow, _ := json.Marshal(jsonthis)
	w.Write(jsonnow)
}
