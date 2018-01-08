package convert

import(
        "github.com/sea350/ustart/server/proto"
        "github.com/sea350/ustart/types"
)

func ToCourseProto(c types.Course) *proto.Course{
    course := proto.Course {
            CourseID: c.CourseID,
            CourseName: c.CourseName,
            BoardIDs: c.BoardIDs,
    }
    return &course
}
