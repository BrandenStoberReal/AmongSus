package main

import "net/http"

func RegisterServerApiHandlers(apiBase string) {
	http.HandleFunc(apiBase+"server/info", serverInfoHandler)
}

func serverInfoHandler(writer http.ResponseWriter, request *http.Request) {

}
