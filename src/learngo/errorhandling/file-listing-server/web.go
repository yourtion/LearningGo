package main

import (
	"learngo/errorhandling/file-listing-server/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {

		// 处理 Panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic:%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handing request: %s\n", err.Error())

			// 处理自定义错误
			if userError, ok := err.(userError); ok {
				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest)
				return
			}

			// 默认处理 Type Assertion
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

// pprof: http://localhost:8888/debug/pprof/
// CPU Profile:  go tool pprof http://localhost:8888/debug/pprof/profile
// more in pprof file

func main() {
	http.HandleFunc("/",
		errWrapper(
			filelisting.HandlerFileLiet))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
