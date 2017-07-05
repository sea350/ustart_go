package main

import(
        _"fmt"
        "net/http"
        "os"
        "log"

        _ "github.com/sea350/ustart/goCode/requests/get"
        "github.com/sea350/ustart/goCode/server/handler"
)

// func handle(w http.ResponseWriter, r *http.Request){
//         fmt.Fprintf(w, "Hi there! I love %s!", r.URL.Path[1:])
// }
//
// func GetBoard(w http.ResponseWriter, r *http.Request){
//         fmt.Fprintf()
// }

const ELASTIC_URL = "http://code.engineering.nyu.edu:9900/"
const USERNAME = "elastic"
const PASSWORD = "elasticpassword"


func main(){
        // Initialize the configuration object
        handler.InitializeConfiguration(ELASTIC_URL, USERNAME, PASSWORD)

        // http.HandleFunc("/", handle)

        http.HandleFunc("/v1/board/get/", handler.GetBoardHandler)
        http.HandleFunc("/v1/board/post/new/", handler.PostNewBoardHandler)
        http.HandleFunc("/v1/board/post/update/", handler.UpdateBoardHandler)

        http.HandleFunc("/v1/user/get/", handler.GetUserHandler)
        http.HandleFunc("/v1/user/post/", handler.IndexUserHandler)
        http.HandleFunc("/v1/user/post/newPrivilege/", handler.AddUserPrivilegeHandler)
        http.HandleFunc("/v1/user/post/modPrivilege/", handler.ModifyExistingUserPrivilegeHandler)
        http.HandleFunc("/v1/user/post/remPrivilege/", handler.RemoveUserPrivilegeHandler)

        http.HandleFunc("/v1/thread/get/", handler.GetThreadHandler)

        http.HandleFunc("/v1/course/get/", handler.GetCourseHandler)
        http.HandleFunc("/v1/course/post/", handler.PostCourseHandler)
        
        args := os.Args
        log.Fatal(http.ListenAndServe(":" + args[1], nil))
}
