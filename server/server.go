package main

import (
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

const magicString string = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./public/handshake.html")
	w.Write([]byte(content))
}

func handShakeHandler(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("Sec-WebSocket-Key")
	accept := calcSecAccept(key)
	log.Println(key, accept)

	h := w.Header()
	h.Set("Connection", "Upgrade")
	h.Set("Upgrade", "websocket")
	h.Set("Sec-WebSocket-Accept", accept)
	w.WriteHeader(http.StatusSwitchingProtocols)
}

func calcSecAccept(key string) string {
	key += magicString
	hasher := sha1.New()
	hasher.Write([]byte(key))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/handshake", handShakeHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
