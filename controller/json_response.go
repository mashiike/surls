package controller

import (
	"encoding/json"
	"net/http"
)

type JSONResponseWriter struct {
	internal http.ResponseWriter
}

func newJSONResponseWriter(w http.ResponseWriter) JSONResponseWriter {
	return JSONResponseWriter{internal: w}
}

func (w JSONResponseWriter) Header() http.Header {
	return w.internal.Header()
}

func (w JSONResponseWriter) WriteHeader(statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.internal.WriteHeader(statusCode)
}

func (w JSONResponseWriter) Write(bs []byte) (int, error) {
	return w.Write(bs)
}

func (w JSONResponseWriter) WriteBody(body interface{}) error {
	return json.NewEncoder(w.internal).Encode(body)
}

func (w JSONResponseWriter) WriteError(statusCode int, err error) error {
	w.WriteHeader(statusCode)
	return w.WriteBody(map[string]interface{}{
		"status":  statusCode,
		"message": http.StatusText(statusCode),
		"detail":  err.Error(),
	})
}
