package get

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    "golang.org/x/net/context"
    "errors"
    "reflect"
    _"encoding/json"
)

const THREAD_INDEX="test-thread_data"
const THREAD_TYPE="thread"

func GetThreadsByThreadID(eclient *elastic.Client, threadID string) ([]types.Thread, error) {
        ctx := context.Background()
        var ret []types.Thread
        exists, err := eclient.IndexExists(THREAD_INDEX).Do(ctx)
        if err != nil {
            return ret, err 
        }
        if !exists {
            return ret, errors.New("Index does not exist")
        }
        termQuery := elastic.NewTermQuery("ThreadID", threadID)
        searchResult, err := eclient.Search().
                Index(THREAD_INDEX).
                Query(termQuery).
                // From(0).Size(MAX_SIZE).
                Do(ctx)
        if err != nil {
                return ret, err 
        }
        var ttyp types.Thread
        for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
                if t, ok := item.(types.Thread); ok {
                        ret = append(ret, t)
                } 
        
        }
        return ret, nil
}
