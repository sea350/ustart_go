package antispam

import "time"

type timelog struct {
	lastTimestamps []time.Time
	lockoutUntil   time.Time
}

type policy struct {
	frequency     int
	withinTime    time.Duration
	lockoutLength time.Duration
}

type spamRecord struct {
	record     map[string]timelog
	spamPolicy policy
}

//spamProtecc is meant to be run when you need generic spam control spam control,
//returns true if you are permitted to continue, false if you should be locked out
func spamProtecc(userID string, trackRecord spamRecord) bool {

	timeSent := time.Now()

	log, exists := trackRecord.record[userID]
	if !exists {
		log.lastTimestamps = append(log.lastTimestamps, timeSent)
		trackRecord.record[userID] = log
		return true
	}

	if len(log.lastTimestamps) >= trackRecord.spamPolicy.frequency {
		if timeSent.Before(log.lockoutUntil) {
			return false
		}
		if log.lastTimestamps[0].Add(trackRecord.spamPolicy.withinTime).Before(timeSent) {
			log.lastTimestamps = log.lastTimestamps[1 : trackRecord.spamPolicy.frequency-1]
			log.lastTimestamps = append(log.lastTimestamps, timeSent)
			trackRecord.record[userID] = log
			return true
		}
		log.lockoutUntil = timeSent.Add(trackRecord.spamPolicy.lockoutLength)
		trackRecord.record[userID] = log
		return false
	}

	return true
}
