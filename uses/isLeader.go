package uses

import (
	elastic "gopkg.in/olivere/elastic.v5"
)

//IsLeader ...
//Returns bool to represent whether member is project leader
//ALso returns index of member
func IsLeader(eclient *elastic.Client, memberID string) (bool, int) {
	var idx int
	for i := range proj.Members {
		idx = i
		if proj.Members[i].MemberID == currentLeaderID && proj.Members[i].Role == 0 {
			return true, i
		}
	}

	return false, idx
}
