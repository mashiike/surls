package controller

import (
	"encoding/json"
	"net/http"
)

type JSONResponseWriter struct {
	http.ResponseWriter
}

func newJSONResponseWriter(w http.ResponseWriter) JSONResponseWriter {
	return JSONResponseWriter{ResponseWriter: w}
}

func (w JSONResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w JSONResponseWriter) WriteHeader(statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w JSONResponseWriter) WriteBody(body interface{}) error {
	return json.NewEncoder(w.ResponseWriter).Encode(body)
}

func (w JSONResponseWriter) WriteError(statusCode int, err error) error {
	w.WriteHeader(statusCode)
	return w.WriteBody(map[string]interface{}{
		"status":  statusCode,
		"message": http.StatusText(statusCode),
		"detail":  err.Error(),
	})
}
