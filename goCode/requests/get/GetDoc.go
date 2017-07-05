package get

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    "golang.org/x/net/context"
    "errors"
    "encoding/json"
)

const DOC_INDEX = "test-documents"
const DOC_TYPE  = "doc"

func GetDocByUniqueID(eclient *elastic.Client, id string) (types.Doc, error){
        ctx := context.Background()
        var ret types.Doc
        exists, err := eclient.IndexExists(DOC_INDEX).Do(ctx)
        if err != nil {
            return ret, err 
        }
        if !exists {
            return ret, errors.New("Index does not exist")
        }
        getResult, err := eclient.Get().
                Index(DOC_INDEX).
                Type(DOC_TYPE).
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
