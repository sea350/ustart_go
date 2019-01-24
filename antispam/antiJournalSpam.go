package antispam

import "time"

//define policy here
var journalRecord = spamRecord{
	record: make(map[string]timelog),
	spamPolicy: policy{
		frequency:     2,
		withinTime:    4 * time.Second,
		lockoutLength: 1 * time.Minute,
	},
}

//AntiJournalSpam is meant to be run when you need journal specific spam control,
//returns true if you are permitted to continue, false if you should be locked out
func AntiJournalSpam(userID string) bool {
	return spamProtecc(userID, journalRecord)
}
