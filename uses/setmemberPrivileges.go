package uses

import (
	types "github.com/sea350/ustart_go/types"
)

//SetMemberPrivileges...
//0 = Creator, 1 = admin, 2 =  member
func SetMemberPrivileges(role int) types.Privileges {
	var memberPrivileges types.Privileges
	switch role {
	case 0:
		memberPrivileges = types.Privileges{
			RoleName:     "Creator",
			RoleID:       role,
			MemberManage: true,
			WidgetManage: true,
			PostManage:   true,
			Icon:         true,
			Banner:       true,
			Links:        true,
			Tags:         true,
		}
	case 1:
		memberPrivileges = types.Privileges{
			// RoleName:     "Member",
			// RoleID:       role,
			// MemberManage: false,
			// WidgetManage: false,
			// PostManage:   false,
			// Icon:         false,
			// Banner:       false,
			// Links:        false,
			// Tags:         false,
			RoleName:     "Member",
			RoleID:       role,
			MemberManage: false,
			WidgetManage: false,
			PostManage:   false,
			Icon:         false,
			Banner:       false,
			Links:        false,
			Tags:         false,
		}
	case 2:
		memberPrivileges = types.Privileges{
			RoleName:     "Admin",
			RoleID:       role,
			MemberManage: false,
			WidgetManage: false,
			PostManage:   true,
			Icon:         false,
			Banner:       false,
			Links:        false,
			Tags:         false,
		}

	}
	return memberPrivileges

}
