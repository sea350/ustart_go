package uses

import (
	"strings"

	types "github.com/sea350/ustart_go/backend/types"
)

//HasEventPrivilege ...
//Returns bool to represent whether member is project leader
//ALso returns index of member
func HasEventPrivilege(privilege string, privs []types.EventPrivileges, member types.EventMembers) bool {

	if member.Role == 0 {
		return true
	}
	privilegeProfile := privs[member.Role]
	var checkPrivilege bool
	switch strings.ToLower(privilege) {
	case "member":
		checkPrivilege = privilegeProfile.MemberManage
	case "widget":
		checkPrivilege = privilegeProfile.WidgetManage
	case "post":
		checkPrivilege = privilegeProfile.PostManage
	case "icon":
		checkPrivilege = privilegeProfile.Icon
	case "banner":
		checkPrivilege = privilegeProfile.Banner
	case "links":
		checkPrivilege = privilegeProfile.Links
	case "tags":
		checkPrivilege = privilegeProfile.Tags
	default:
		return false
	}

	return checkPrivilege
}
