package get

import (
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//IsEventGuest ...
func IsEventGuest(eclient *elastic.Client, guestID string, event types.Events) bool {
	if len(event.Guests) < 1 {
		return false
	}
	/*
		for gust, _ := range event.Guests {
			if event.Guests[gust].GuestID == guestID {
				return true
			}
		}
	*/

	return false

}
