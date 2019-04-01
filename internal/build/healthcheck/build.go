package healthcheck

import "net/http"

// Build uses the configuraiotn to return a http server for healthchecks
func Build(config *Config) *http.Server {
	var (
		mux    = http.NewServeMux()
		server = &http.Server{
			Addr: config.Addr,
		}
	)

	mux.HandleFunc(config.Path, func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("healthy"))
	})

	server.Handler = mux

	return server
}
