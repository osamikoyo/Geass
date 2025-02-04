package transport

import "net/http"

type handler struct {}

func (h handler) MainHandler(w http.ResponseWriter, r *http.Request) error {
	
}