package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse はJSONレスポンスを送信します
func JSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ParseJSONBody はリクエストボディからJSONを解析します
func ParseJSONBody(r *http.Request, dest interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(dest)
}
