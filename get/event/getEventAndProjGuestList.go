package get

import (
	"errors"

	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EventAndProjGuestList ...
func EventAndProjGuestList(eclient *elastic.Client, eventID string) (types.Events, []types.EventProjectGuests, error) {
	evnt, err := EventByID(eclient, eventID)
	if err != nil {
		return types.Events{}, []types.EventProjectGuests{}, err
	}

	if len(evnt.ProjectGuests) < 1 {
		return types.Events{}, []types.EventProjectGuests{}, errors.New("Event has zero Project Guests")
	}

	return evnt, evnt.ProjectGuests, err
}
