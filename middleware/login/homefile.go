package login

import (
  "net/http"
  elastic "gopkg.in/olivere/elastic.v5"
  "github.com/gorilla/sessions"
  client "github.com/sea350/ustart_go/middleware/clientstruct"

)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code 


func Home (w http.ResponseWriter, r *http.Request){
  session, _ := store.Get(r, "session_please")
  test1, _ := session.Values["DocID"]
  if (test1 != nil){
    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound) }
  session.Save(r, w)
  cs := client.ClientSide{}
  //fmt.Println("helllo")
  client.RenderTemplate(w,"templateNoUser2",cs)
  client.RenderTemplate(w,"nil-index2",cs)
  client.RenderTemplate(w,"template-footer-nil",cs)
}


