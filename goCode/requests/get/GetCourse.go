package get

import(
    elastic "gopkg.in/olivere/elastic.v5"
    types "github.com/nicolaifsf/TheTandon-Server/types"
    "golang.org/x/net/context"
    "errors"
    "encoding/json"
)

const COURSE_INDEX = "test-course_data"
const COURSE_TYPE  = "course"

func GetCourseByUniqueID(eclient *elastic.Client, id string) (types.Course, error){
        ctx := context.Background()
        var ret types.Course
        exists, err := eclient.IndexExists(COURSE_INDEX).Do(ctx)
        if err != nil {
            return ret, err 
        }
        if !exists {
            return ret, errors.New("Index does not exist")
        }
        getResult, err := eclient.Get().
                Index(COURSE_INDEX).
                Type(COURSE_TYPE).
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
