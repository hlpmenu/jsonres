package jsonres

import "net/http"

// ServeUserUnauthorized serves a simple status response with just status and message fields.
// Use this for success/error responses where no data payload is needed.
// Serves: {"status": "error", "message": "unauthorized"}
// Status: 401
func ServeUnauthorized(w http.ResponseWriter) error {
	return ServeJsonStatusRes("error", "unauthorized", http.StatusUnauthorized, w)
}

// ServeInternalServerError serves a simple status response with just status and message fields.
// Use this for success/error responses where no data payload is needed.
// Serves: {"status": "error", "message": "internal server error"}
// Status: 500
func ServeInternalServerError(w http.ResponseWriter) error {
	return ServeJsonStatusRes("error", "internal server error", http.StatusInternalServerError, w)
}

// ServeNotFound serves a simple status response with just status and message fields.
// Use this for success/error responses where no data payload is needed.
// Serves: {"status": "error", "message": "not found"}
// Status: 404
func ServeNotFound(w http.ResponseWriter) error {
	return ServeJsonStatusRes("error", "not found", http.StatusNotFound, w)
}

// ServeBadRequest serves a simple status response with just status and message fields.
// Use this for success/error responses where no data payload is needed.
// Serves: {"status": "error", "message": "bad request"}
// Status: 400
func ServeBadRequest(w http.ResponseWriter) error {
	return ServeJsonStatusRes("error", "bad request", http.StatusBadRequest, w)
}

// ServeForbidden serves a simple status response with just status and message fields.
// Use this for success/error responses where no data payload is needed.
// Serves: {"status": "error", "message": "forbidden"}
// Status: 403
func ServeForbidden(w http.ResponseWriter) error {
	return ServeJsonStatusRes("error", "forbidden", http.StatusForbidden, w)
}

// ServeOk serves a simple status response with just status and message fields.
// Use this for success/error responses where no data payload is needed.
// Serves: {"status": "ok", "message": "ok"}
// Status: 200
func ServeSimpleOk(w http.ResponseWriter) {
	ServeJsonStatusRes("ok", "ok", http.StatusOK, w)
}

func Example(w http.ResponseWriter, status, message string) error {
	return ServeJsonStatusRes(status, message, http.StatusUnauthorized, w)
}
