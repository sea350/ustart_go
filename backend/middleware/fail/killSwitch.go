package fail

import (
	"net/http"
	"os"
)

//KillSwitch ... kills process
func KillSwitch(w http.ResponseWriter, r *http.Request) {
	os.Exit(420)
}
