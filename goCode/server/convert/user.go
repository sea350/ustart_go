package convert

import(
        "github.com/nicolaifsf/TheTandon-Server/server/proto"
        "github.com/nicolaifsf/TheTandon-Server/types"
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
