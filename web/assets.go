package web

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func AddHealth(m *http.ServeMux) { m.HandleFunc("/health", Health) }
