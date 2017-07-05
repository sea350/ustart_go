package get

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    "golang.org/x/net/context"
    "errors"
    "reflect"
    "encoding/json"
)


const USER_INDEX="test-user_data"
const USER_TYPE="user"

func GetUsersByUserID(eclient *elastic.Client, userID string) ([]types.User, error) {
        ctx := context.Background()
        var ret []types.User
        exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
        if err != nil {
            return ret, err 
        }
        if !exists {
            return ret, errors.New("Index does not exist")
        }
        termQuery := elastic.NewTermQuery("UserID", userID)
        searchResult, err := eclient.Search().
                Index(USER_INDEX).
                Query(termQuery).
                // From(0).Size(MAX_SIZE).
                Do(ctx)
        if err != nil {
                return ret, err 
        }
        var ttyp types.User
        for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
                if t, ok := item.(types.User); ok {
                        ret = append(ret, t)
                } 
        
        }
        return ret, nil
}

func GetUserByUniqueID(eclient *elastic.Client, id string) (types.User, error){
        ctx := context.Background()
        var ret types.User
        exists, err := eclient.IndexExists(USER_INDEX).Do(ctx)
        if err != nil {
            return ret, err 
        }
        if !exists {
            return ret, errors.New("Index does not exist")
        }
        getResult, err := eclient.Get().
                Index(USER_INDEX).
                Type(USER_TYPE).
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
