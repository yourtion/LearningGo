package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter,
	resuest *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	resuest *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter,
	resuest *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter,
	resuest *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter,
	resuest *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	resuest *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func verifyResponse(resp *http.Response,
	expectedCode int,
	expectedMessage string,
	t *testing.T) {

	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode ||
		body != expectedMessage {
		t.Errorf("except (%d, %s); got (%d, %s)",
			expectedCode, expectedMessage,
			resp.StatusCode, body)
	}
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://blog.yourtion.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(
			http.HandlerFunc(f))

		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}
