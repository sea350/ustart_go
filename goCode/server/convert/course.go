package convert

import(
        "github.com/nicolaifsf/TheTandon-Server/server/proto"
        "github.com/nicolaifsf/TheTandon-Server/types"
)

func ToCourseProto(c types.Course) *proto.Course{
    course := proto.Course {
            CourseID: c.CourseID,
            CourseName: c.CourseName,
            BoardIDs: c.BoardIDs,
    }
    return &course
}
