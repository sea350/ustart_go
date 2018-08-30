package event

import (
	"encoding/json"
	"net/http"
)

//FindEventMember ... find event members
func FindEventMember(w http.ResponseWriter, r *http.Request) {
	term := r.FormValue("term")
	jsonthis := []string{"Magic Johnson", "Magic White", "Magic Pomper", term}
	jsonnow, _ := json.Marshal(jsonthis)
	w.Write(jsonnow)
}

//FindEventGuest ... find event guest
func FindEventGuest(w http.ResponseWriter, r *http.Request) {
	term := r.FormValue("term")
	jsonthis := []string{"Magic Johnson", "Magic White", "Magic Pomper", term}
	jsonnow, _ := json.Marshal(jsonthis)
	w.Write(jsonnow)
}

//FindEventProject ... find event project
func FindEventProject(w http.ResponseWriter, r *http.Request) {
	term := r.FormValue("term")
	jsonthis := []string{"Magic Johnson", "Magic White", "Magic Pomper", term}
	jsonnow, _ := json.Marshal(jsonthis)
	w.Write(jsonnow)
}
