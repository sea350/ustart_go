package convert

import(
        "github.com/nicolaifsf/TheTandon-Server/server/proto"
        "github.com/nicolaifsf/TheTandon-Server/types"
)

func ToPrivilegeProto(p types.Privilege) *proto.Privilege{
        privilege := proto.Privilege {
                CourseName: p.CourseName,
                CourseID: p.CourseID,
                AccessLevel: p.AccessLevel,
        }
        return &privilege
}
