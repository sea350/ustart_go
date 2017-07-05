package post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
)

func PostComment(eclient *elastic.Client, thread types.Thread, comment types.Comment) error {
    //  First, we append to the containing thread's comments array.
    thread.Comments = append(thread.Comments, comment)
    //  Finally, we push up the thread as a new version
    return UpdateThread(eclient, thread)
}
