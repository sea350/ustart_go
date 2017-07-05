package post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    "github.com/nicolaifsf/TheTandon-Server/server/convert"
    "golang.org/x/net/context"
    "errors"
)

const DOC_INDEX = "test-documents"
const DOC_TYPE  = "doc"

func PostDoc(eclient *elastic.Client, document types.Doc, board types.Board) error {
    ctx := context.Background()
    exists, err := eclient.IndexExists(DOC_INDEX).Do(ctx)
    if err != nil {
        return err 
    }
    if !exists {
        return errors.New("Index does not exist")
    }
    //	First, we have to push the Document into the index
    _, err = eclient.Index().
        Index(DOC_INDEX).
        Type(DOC_TYPE).
        BodyJson(document).
        Do(ctx)
    if err != nil {
    	return err 
    }
    //	Next, we update the containing board's docIDs array.
    board.DocIDs = append(board.DocIDs, document.DocID)

    //	Finally, we push our updated copy of the Board into the index.
    return UpdateBoard(eclient, *convert.ToBoardProto(board))   
}
