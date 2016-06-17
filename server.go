package main

import (
	"net/http"
	"encoding/json"	
	"fmt"
	"log"
	"browserstats/internal"
)

const Port = "8080"
const FromParamKey = "from"
const ToParamKey = "to"

var	statCounterClient = internal.NewStatCounterClient()	
var	requestValidator = internal.NewRequestValidator()
func main() {
	http.HandleFunc("/stats", getStats)
	fmt.Println("Listing on port", Port)
	log.Fatal(http.ListenAndServe(":" + Port, nil))
}

func getStats(response http.ResponseWriter, request *http.Request) {
	statCounterChannel := make(chan *internal.StatCounterResponse)
	fromDate, toDate, err := requestValidator.Validate(request.URL.Query())

	if (err != nil) {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	go statCounterClient.GetBrowserShare(fromDate, toDate, statCounterChannel)

	statCounterResponse := <- statCounterChannel

	if (statCounterResponse.Error != nil) {
		http.Error(response, statCounterResponse.Error.Error(), http.StatusInternalServerError)
		return	
	}

	responseData := internal.NewResponse(statCounterResponse.Response)
	jsonBytes, marshalErr := json.Marshal(&responseData)

	if (marshalErr != nil) {
		http.Error(response, statCounterResponse.Error.Error(), http.StatusInternalServerError)
		return	
	}

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonBytes)
}