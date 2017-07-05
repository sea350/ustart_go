package get

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    "golang.org/x/net/context"
    "errors"
    "encoding/json"
)

const BOARD_INDEX= "test-board_data"
const BOARD_TYPE = "board"

func GetBoardByUniqueID(eclient *elastic.Client, id string) (types.Board, error){
        ctx := context.Background()
        var ret types.Board
        exists, err := eclient.IndexExists(BOARD_INDEX).Do(ctx)
        if err != nil {
            return ret, err 
        }
        if !exists {
            return ret, errors.New("Index does not exist")
        }
        getResult, err := eclient.Get().
                Index(BOARD_INDEX).
                Type(BOARD_TYPE).
                Id(id).
                Do(ctx)
        if err != nil {
                return ret, err 
        }
        err = json.Unmarshal(*getResult.Source, &ret)
        if err != nil {
            return ret, err 
        }
        return ret, nil
}
