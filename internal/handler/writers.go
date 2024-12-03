package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendJSON(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)

	w.Header().Set(HeaderContentType, JsonContentType)
	sendBytes(w, code, res)
}

func sendBytes(w http.ResponseWriter, code int, data []byte) {
	w.WriteHeader(code)
	if _, err := w.Write(data); err != nil {
		fmt.Println("write manager couriers report result failed: ", err)
	}
}
