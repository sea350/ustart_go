package convert

import(
        "github.com/sea350/ustart/server/proto"
        "github.com/sea350/ustart/types"
)

func ToUserProto(u types.User) *proto.User {
        var privs []*proto.Privilege
        for _, element := range u.Privileges {
            privs = append(privs, ToPrivilegeProto(element));
        }

        user := proto.User {
            UserID: u.UserID,
            Username: u.Username,
            Password: u.Password,
            Privileges: privs,
        }
        return &user
}
