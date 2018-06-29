package chat

import (
	template "text/template"
	"net/http"
)

//PageHandler ... draws chat page
func PageHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("/ustart/ustart_front/cuzsteventoldmeto.html")).Execute(w, r)
	})
}
