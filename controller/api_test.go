package controller_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mashiike/surls/controller"
	"github.com/mashiike/surls/usecase"
	"github.com/mashiike/surls/usecase/stub"
	"github.com/stretchr/testify/assert"
)

func TestAPIController(t *testing.T) {
	c := controller.NewAPIController(&usecase.Usecase{
		CreateShortcutBoundary: stub.NewCreateShortcutInteractor(),
	})
	router := mux.NewRouter()
	c.RegisterRoute(router)

	cases := []struct {
		casename string
		method   string
		target   string
		reqBody  string
		status   int
		respBody string
	}{
		{
			casename: "create success",
			method:   "POST",
			target:   "/urls",
			reqBody:  `{"long_url":"https://example.org"}`,
			status:   http.StatusOK,
			respBody: fmt.Sprintf(`{"shortcut_id":"%s","long_url":"https://example.org"}`, stub.DummyShortcutID),
		},
		{
			casename: "create internal server error",
			method:   "POST",
			target:   "/urls",
			reqBody:  fmt.Sprintf(`{"long_url":"%s"}`, stub.FaildLongURL),
			status:   http.StatusInternalServerError,
			respBody: fmt.Sprintf(`{"status":500,"message":"Internal Server Error","detail": "stub failed url"}`),
		},
		{
			casename: "create bad request",
			method:   "POST",
			target:   "/urls",
			reqBody:  `hoge`,
			status:   http.StatusBadRequest,
			respBody: fmt.Sprintf(`{"status":400,"message":"Bad Request","detail": "invalid character 'h' looking for beginning of value"}`),
		},
	}

	for _, c := range cases {
		t.Run(c.casename, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(c.method, c.target, bytes.NewBufferString(c.reqBody))
			router.ServeHTTP(w, req)
			assert.Equalf(t, c.status, w.Code, "expected HTTP Status Code %d", c.status)
			assert.JSONEq(t, c.respBody, w.Body.String(), "HTTP Response Body")
		})
	}
}
