package jsonres

import (
	"net/http"
)

// JsonStatusRes returns a JSON status response with just status and message fields.
// Use this for success/error responses where no data payload is needed.
// Format: {"status": string, "message": string}
func JsonStatusRes(status, message string) ([]byte, error) {
	res, err := JsonStatusResponse{
		Status:  status,
		Message: message,
	}.ToJson()
	if err != nil {
		return nil, err
	}
	return res, nil

}

// ServeJsonStatusRes serves a simple status response with just status and message fields.
// Use this for success/error responses where no data payload is needed.
// Format: {"status": string, "message": string}
func ServeJsonStatusRes(status, message string, statuscode int, w http.ResponseWriter) error {
	res, err := JsonStatusRes(status, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	_, err = w.Write(res)
	return err
}
