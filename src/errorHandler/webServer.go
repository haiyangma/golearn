package main

import (
	"errorHandler/fileListServer"
	"net/http"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover();r!=nil{
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()
		err := handler(writer,request)
		if err!=nil{
			log.Warn("Error handling request %s",err.Error())
			if userError,ok := err.(userError);!ok{
				http.Error(
						writer,
						userError.Message(),
						http.StatusBadRequest,
					)
			}
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError

			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

type userError interface {
	error
	Message() string
}


func main() {
	http.HandleFunc("/list/", errWrapper(fileListServer.HandleFileList))

	err := http.ListenAndServe(":8888",nil)
	if err !=nil{
		panic(err)
	}

}
