package httpx

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, status int, message string, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": message,
		"code":    code,
	},
	)
}
