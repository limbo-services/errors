package errors

import (
	"encoding/json"
	"io"
	"net/http"
)

// HTTPError allows an error to be used as an HTTP response
type HTTPError interface {
	error
	HTTPStatusCode() int
	HTTPMessage() string
}

// WriteHTTPError writes a JSON representation of err
func WriteHTTPError(w http.ResponseWriter, r *http.Request, err error) error {
	if x, ok := err.(HTTPError); ok && x != nil {
		var (
			code    = x.HTTPStatusCode()
			payload = struct {
				Error string `json:"error"`
			}{
				Error: x.HTTPMessage(),
			}
		)

		data, err := json.Marshal(&payload)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(code)
		w.Write(data)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, `{"error": "internal server error"}`)
	return nil
}
