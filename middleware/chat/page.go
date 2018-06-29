package chat

import (
	"html/template"
	"net/http"
)

//Page ... draws chat page
func Page(tpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r)
	})
}
