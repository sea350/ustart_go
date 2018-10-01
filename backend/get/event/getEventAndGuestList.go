package get

import (
	"errors"

	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventAndGuestList ...
func EventAndGuestList(eclient *elastic.Client, eventID string) (types.Events, []types.EventGuests, error) {
	evnt, err := EventByID(eclient, eventID)
	if err != nil {
		return types.Events{}, []types.EventGuests{}, err
	}

	if len(evnt.Guests) < 1 {
		return types.Events{}, []types.EventGuests{}, errors.New("Event has zero guests")
	}

	return evnt, evnt.Guests, err
}
