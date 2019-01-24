package antispam

import "time"

//define policy here
var chatRecord = &spamRecord{
	record: make(map[string]timelog),
	spamPolicy: policy{
		frequency:     5,
		withinTime:    1 * time.Second,
		lockoutLength: 30 * time.Second,
	},
}

//AntiChatSpam is meant to be run when you need chat specific spam control spam control,
//returns true if you are permitted to continue, false if you should be locked out
func AntiChatSpam(userID string) bool {
	return spamProtecc(userID, chatRecord)
}
