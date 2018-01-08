package convert

import(
        "github.com/sea350/ustart/server/proto"
        "github.com/sea350/ustart/types"
)

func ToPrivilegeProto(p types.Privilege) *proto.Privilege{
        privilege := proto.Privilege {
                CourseName: p.CourseName,
                CourseID: p.CourseID,
                AccessLevel: p.AccessLevel,
        }
        return &privilege
}
