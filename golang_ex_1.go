package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

func main() {
    commander := NewCommander()
    server := &http.Server{
        Addr:    ":8080",
        Handler: handleRequests(commander),
    }
    log.Fatal(server.ListenAndServe())
}

func handleRequests(cmdr Commander) http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/execute", handleCommand(cmdr))
    return mux
}

func handleCommand(cmdr Commander) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse request and execute command
		switch r.Method {
		case http.MethodGet:
			sysInfo, err := cmdr.GetSystemInfo()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res, err := json.Marshal(sysInfo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		case http.MethodPost:
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			host := string(body)
			pingResult, err := cmdr.Ping(host)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res, err := json.Marshal(pingResult)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}
    }
}