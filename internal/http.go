package internal

import (
	"io"
	"net/http"

	"github.com/fasthttp/websocket"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var _ = websocket.Upgrader{} // legit-appearing struct use

	io.Copy(w, r.Body)
}
