package post

import(
    elastic "gopkg.in/olivere/elastic.v5"
    // types "github.com/nicolaifsf/TheTandon-Server/types"
    "github.com/nicolaifsf/TheTandon-Server/server/proto"
    "golang.org/x/net/context"
    "errors"
)

const COURSE_INDEX = "test-course_data"
const COURSE_TYPE  = "course"

//  It is assumed that the course argument is empty.  
//  If one attempts to add a course with board-children, 
//      the accompanying children will *not* be added.  
//      Any desired children must be added in separate calls.

func PostCourse(eclient *elastic.Client, course proto.Course) error {
    ctx := context.Background()
    exists, err := eclient.IndexExists(COURSE_INDEX).Do(ctx)
    if err != nil {
        return err 
    }
    if !exists {
        return errors.New("Index does not exist")
    }
    _, err = eclient.Index().
            Index(COURSE_INDEX).
            Type(COURSE_TYPE).
            BodyJson(course).
            Do(ctx)

    if err != nil {
    	return err 
    }
    return nil
}
