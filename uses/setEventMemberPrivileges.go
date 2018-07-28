package uses

import (
	types "github.com/sea350/ustart_go/types"
)

//SetEventMemberPrivileges ...
//0 = Creator, 1 = admin, 2 = member
func SetEventMemberPrivileges(role int) types.EventPrivileges {
	var eventmemPrivileges types.EventPrivileges
	switch role {
	case 0:
		eventmemPrivileges = types.EventPrivileges{
			RoleName:     "Creator",
			RoleID:       role,
			MemberManage: true,
			WidgetManage: true,
			PostManage:   true,
			Banner:       true,
			Links:        true,
			Tags:         true,
		}
	case 1:
		eventmemPrivileges = types.EventPrivileges{
			RoleName:     "Admin",
			RoleID:       role,
			MemberManage: false,
			WidgetManage: true,
			PostManage:   true,
			Icon:         false,
			Banner:       false,
			Links:        false,
			Tags:         false,
		}
	case 2:
		eventmemPrivileges = types.EventPrivileges{
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
	}
	return eventmemPrivileges
}
