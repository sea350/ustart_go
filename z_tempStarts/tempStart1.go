package main

import (
	"database/sql"
	htype "html/template"
	"log"
	"net/http"

	"github.com/lib/pq"
	"github.com/sea350/ustart_go/middleware/fail"
)

var livePort = "5000"
var templates = htype.Must(htype.ParseFiles("/ustart/ustart_front/nil-index2.html",
	"/ustart/ustart_front/index.php",
	"/ustart/ustart_front/template2-nil.html",
	"/ustart/ustart_front/index1.html"))

func main() {
	/*
		Lines 18-19 handle the static file locating
		If we wanted to reorganize file/folder locations, this is one of 3 things that would have to change
		In executeTemplates you will need to make the same changes
		The other being the relative link on the actual html pages
	*/
	// fs := http.FileServer(http.Dir("/home/rr2396/www/"))
	_, _ = http.Get("http://ustart.today:" + livePort + "/KillUstartPlsNoUserinoCappucinoDeniro")
	fs := http.FileServer(http.Dir("/ustart/ustart_front/"))

	http.Handle("/ustart_front/", http.StripPrefix("/ustart_front/", fs))
	/*
		The following are all the handlers we have so fart.
	*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")

		if email != `` {
			conn := "host= ustart.today port=5432 dbname=ustart user=ustart password=~m3lanKollymemes"
			_ = pq.Efatal
			db, err := sql.Open("postgresql", conn)
			defer db.Close()
			if err != nil {
				log.Println(err)
			} else {
				_, err := db.Exec("insert into newsletter (uname, email) values ('" + name + "', '" + email + "')")
				if err != nil {
					log.Println(err)
				}
			}
		}

		err := templates.ExecuteTemplate(w, "index1.html", nil)
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/404/", fail.Fail)
	http.HandleFunc("/KillUstartPlsNoUserinoCappucinoDeniro", fail.KillSwitch)

	http.ListenAndServe(":"+livePort, nil)
}
