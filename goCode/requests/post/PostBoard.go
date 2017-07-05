package post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    _"fmt"
    "golang.org/x/net/context"
    "github.com/nicolaifsf/TheTandon-Server/server/proto"
    "errors"
    _"encoding/json"
)

const BOARD_INDEX = "test-board_data"
const BOARD_TYPE = "board"

//  It is assumed that the board struct is empty.  
//  If one attempts to add a board with thread-children, 
//      the accompanying children will *not* be added automatically.  
//      These children must be added in separate calls.

func PostNewBoard(eclient *elastic.Client, board proto.Board, course types.Course) error {
    return nil // TODO: FIX THIS!!!
    // ctx := context.Background()
    // exists, err := eclient.IndexExists(BOARD_INDEX).Do(ctx)
    // if err != nil {
    //     return err 
    // }
    // if !exists {
    //     return errors.New("Index does not exist")
    // }
    //
    // asJSON, err := json.Marshal(board)
    // if err != nil {
    //     return errors.New("Error marshalling to JSON") 
    // }
    // fmt.Println(asJSON)
    //
    // //  First, we post the new board.
    // _, err = eclient.Index().
    //     Index(BOARD_INDEX).
    //     Type(BOARD_TYPE).
    //     BodyJson(asJSON).
    //     Do(ctx)
    // if err != nil {
    //     return err 
    // }
    // //  Next, we append to the containing course's boardIDs array.
    // .BoardIDs = append(course.BoardIDs, board.BoardID)
    //
    // //  Finally, we push up the course as a new version.
    // //  Note:   PostCourse is the same as a hypothetical UpdateCourse,
    // //          since Course has no parents.  However, PostNewBoard
    // //          differs from UpdateBoard, as Board does have parents.
    // return PostCourse(eclient, board)
}

func UpdateBoard(eclient *elastic.Client, updatedBoard proto.Board) error {
    ctx := context.Background()
    exists, err := eclient.IndexExists(BOARD_INDEX).Do(ctx)
    if err != nil {
        return err 
    }
    if !exists {
        return errors.New("Index does not exist")
    }
    _, err = eclient.Index().
        Index(BOARD_INDEX).
        Type(BOARD_TYPE).
        BodyJson(updatedBoard).
        Do(ctx)
    if err != nil {
        return err 
    }
    return nil
}
