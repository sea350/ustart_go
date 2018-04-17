package settings

type response struct {
	Successful bool   `json:"Successful"`
	ErrMsg     error  `json:"ErrMsg"`
	Retreived  string `json:"Retreived"`
	//ColorPalette string `json:"ColorPalette"`
}

func (r *response) update(s bool, e error, ret string) {
	r.Successful = s
	r.ErrMsg = e

	if e != nil {
		panic(e)
	}
}
