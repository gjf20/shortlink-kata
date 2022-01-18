package shortlink

import "net/http"

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", createHandler)

	return mux
}
