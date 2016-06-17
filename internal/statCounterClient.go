package internal

import (
	"bytes"
	"encoding/xml"	
	"fmt"
	"io"
	"net/http"
)

type StatCounterClient struct {
	BaseUrl string
	HttpClient *http.Client
}

func NewStatCounterClient() (*StatCounterClient) {
	client := new(StatCounterClient)
	client.HttpClient = new(http.Client);
	client.BaseUrl = "http://gs.statcounter.com/chart.php?bar=1&device=Desktop%2C%20Tablet%20%26%20Console&device_hidden=desktop%2Btablet%2Bconsole&multi-device=true&statType_hidden=browser_version&region_hidden=ww&granularity=monthly&statType=Browser%20Version&region=Worldwide"

	return client
}

func (statCounterClient *StatCounterClient) GetBrowserShare(fromDate string, toDate string, channel chan *StatCounterResponse) {
	url := fmt.Sprintf("%v&fromMonthYear=%v&toMonthYear=%v", statCounterClient.BaseUrl, fromDate, toDate)
	response, err := statCounterClient.HttpClient.Get(url)

	if err != nil { channel <- NewStatCounterResponse(nil, err) }

	statCounterChart, deserialiseErr := statCounterClient.Deserialise(response.Body)

	channel <- NewStatCounterResponse(statCounterChart, deserialiseErr)
}

func (statCounterClient *StatCounterClient) Deserialise(body io.ReadCloser) (*StatCounterChart, error) {
	var statCounterChart *StatCounterChart

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(body)
	err := xml.Unmarshal(buffer.Bytes(), &statCounterChart)
	body.Close()

	return statCounterChart, err
}