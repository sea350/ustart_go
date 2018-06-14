package uses

import (
	"strings"

	types "github.com/sea350/ustart_go/types"
)

//HasPrivilege ...
//Returns bool to represent whether member is project leader
//ALso returns index of member
func HasPrivilege(privilege string, proj types.Project, member types.Member) bool {

	privilegeProfile := proj.PrivilegeProfiles[member.Role]
	var checkPrivilege string
	switch strings.ToLower(privilege) {
	case "creator":
		return true
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

	return checkPrivilege == true

}
