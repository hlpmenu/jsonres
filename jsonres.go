package jsonres

import (
	"net/http"

	json "github.com/goccy/go-json"
)

type JsonStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (res JsonStatusResponse) Serve(w http.ResponseWriter, statuscode int) {
	jsonb, err := res.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	w.Write(jsonb)

}

func (res JsonStatusResponse) ToJson() ([]byte, error) {
	jsonb, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return jsonb, nil
}
