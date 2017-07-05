package post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    "github.com/nicolaifsf/TheTandon-Server/server/convert"
    "golang.org/x/net/context"
    "errors"
)

const THREAD_INDEX = "test-thread_data"
const THREAD_TYPE  = "thread"

func PostNewThread(eclient *elastic.Client, thread types.Thread, board types.Board) error {
	ctx := context.Background()
    exists, err := eclient.IndexExists(THREAD_INDEX).Do(ctx)
    if err != nil {
        return err 
    }
    if !exists {
        return errors.New("Index does not exist")
    }
    //  First, we post the new thread.
    _, err = eclient.Index().
        Index(THREAD_INDEX).
        Type(THREAD_TYPE).
        BodyJson(thread).
        Do(ctx)
    if err != nil {
        return err 
    }
    //  Next, we append to the containing board's threadIDs array.
    board.ThreadIDs = append(board.ThreadIDs, thread.ThreadID)

    //  Finally, we push up the Board as a new version
    return UpdateBoard(eclient, *convert.ToBoardProto(board))
}

//  Assumes that the thread is already indexed in the parent Board.
func UpdateThread(eclient *elastic.Client, updatedThread types.Thread) error {
    ctx := context.Background()
    exists, err := eclient.IndexExists(THREAD_INDEX).Do(ctx)
    if err != nil {
        return err 
    }
    if !exists {
        return errors.New("Index does not exist")
    }
    _, err = eclient.Index().
        Index(THREAD_INDEX).
        Type(THREAD_TYPE).
        BodyJson(updatedThread).
        Do(ctx)
    if err != nil {
        return err 
    }
    return nil
}
