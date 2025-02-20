package jsonres

import (
	"errors"
	"net/http"

	json "github.com/goccy/go-json"
)

var (
	ErrNilObject     = errors.New("object is nil")
	ErrMarshalFailed = errors.New("json marshal failed")
)

type JsonApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func IsNilObject(err error) bool {
	if err == nil {
		return false
	}
	return err == ErrNilObject
}

// ServeJsonBytes writes raw JSON bytes to the response writer with the specified status code.
// Use this when you already have marshalled JSON data.
func ServeJsonBytes(jsonb []byte, statuscode int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	_, err := w.Write(jsonb)
	return err
}

// MarshalAndServe marshals any object to JSON and serves it directly.
// Use this when you want to send a raw object as JSON response.
// The object will be served as-is, without wrapping it in an API response structure.
// Returns error if marshalling fails.
func MarshalAndServe(obj interface{}, statuscode int, w http.ResponseWriter) error {
	if obj == nil {
		w.WriteHeader(500)
		//nolint:errcheck
		w.Write([]byte(`{"status:"error","message":"internal server error", "error": "internal server error"}`))
		return ErrNilObject
	}
	jsonb, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(500)
		//nolint:errcheck
		w.Write([]byte(`{"status:"error","message":"internal server error", "error": "internal server error"}`))
		return ErrMarshalFailed
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	_, err = w.Write(jsonb)
	return err
}

// MarshallApiResponse marshals an object and wraps it in a standard API response format:
//
//	{
//	  "status": string,
//	  "message": string,
//	  "data": object
//	}
//
// Use this when you need the response wrapped in the standard API response structure.
// Returns error if marshalling fails.
//
//nolint:var-naming
func MarshallApiResponse(status string, message string, obj interface{}, statuscode int, w http.ResponseWriter) error {
	if obj == nil {
		w.WriteHeader(http.StatusInternalServerError)
		//nolint:errcheck
		w.Write([]byte(`{"status:"error","message":"internal server error", "error": "internal server error"}`))
		return ErrNilObject
	}
	if status == "" {
		status = "success"
	}
	if message == "" {
		message = "success"
	}
	resObj := JsonApiResponse{
		Status:  status,
		Message: message,
		Data:    obj,
	}
	jsonb, err := json.Marshal(resObj)
	if err != nil {
		w.WriteHeader(500)
		//nolint:errcheck
		w.Write([]byte(`{"status:"error","message":"internal server error", "error": "internal server error"}`))
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	_, err = w.Write(jsonb)
	return err
}
