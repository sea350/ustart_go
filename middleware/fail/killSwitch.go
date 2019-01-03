package fail

import (
	"net/http"
	
)

//KillSwitch ... kills process
func KillSwitch(w http.ResponseWriter, r *http.Request) {
	os.Exit(420)
}
