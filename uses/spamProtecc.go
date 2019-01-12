package uses

import "time"

//define policy here
const (
	frequency         = 5
	withinTimeSeconds = 1
	lockoutSeconds    = 30
)

var lastTransacts = make(map[string]timelog)

type timelog struct {
	lastTimestamps []time.Time
	lockoutUntil   time.Time
}

//SpamProtecc is meant to be run when you need spam control,
//returns true if you are permitted to continue, false if you should be locked out
func SpamProtecc(userID string) bool {

	timeSent := time.Now()

	log, exists := lastTransacts[userID]
	if !exists {
		log.lastTimestamps = append(log.lastTimestamps, timeSent)
		lastTransacts[userID] = log
		return true
	}

	if len(log.lastTimestamps) >= frequency {
		if timeSent.Before(log.lockoutUntil) {
			return false
		}
		if log.lastTimestamps[0].Add(withinTimeSeconds * time.Second).Before(timeSent) {
			log.lastTimestamps = log.lastTimestamps[1 : frequency-1]
			log.lastTimestamps = append(log.lastTimestamps, timeSent)
			lastTransacts[userID] = log
			return true
		} else {
			log.lockoutUntil = timeSent.Add(lockoutSeconds * time.Second)
			lastTransacts[userID] = log
			return false
		}
	}

	return true
}

//SpamRemove needs to be executed upon logout to save on RAM
func SpamRemove(userID string) {

	delete(lastTransacts, userID)
}
