package web

import "net/http"

// Server serves all of ustart's frontend assets to the browser
type Server struct {
	port      string
	assetRoot string
}

// New returns a new Server, given the passed in config
func New(cfg *Config) (*Server, error) {
	return &Server{
		port:      cfg.Port,
		assetRoot: cfg.AssetsRoot,
	}, nil
}

// Run starts the Server
func (ws *Server) Run() error {
	_, _ = http.Get("http://ustart.today:" + ws.port + "/KillUstartPlsNoUserinoCappucinoDeniro")
	fs := http.FileServer(http.Dir("/ustart/ustart_front/"))
	http.Handle("/ustart_front/", http.StripPrefix("/ustart_front/", fs))

	return http.ListenAndServe(":"+ws.port, nil)
}
