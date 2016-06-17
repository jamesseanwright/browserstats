package main

import (
	"net/http"
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
	//statCounterChannel := make(chan *internal.StatCounterResponse)
	fromDate, toDate, err := requestValidator.Validate(request.URL.Query())

	if (err != nil) {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	statCounterClient.GetBrowserShare(fromDate, toDate)
}