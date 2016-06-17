package internal

import (
	"bytes"
	"encoding/xml"	
	"fmt"
	"io"
	"net/http"
	"github.com/hashicorp/golang-lru"
)

const CacheSize = 500

var cache, err = lru.New(CacheSize)

type StatCounterClient struct {
	BaseUrl string
	HttpClient *http.Client
	Cache *lru.Cache
}

func NewStatCounterClient() (*StatCounterClient) {
	cache, err := lru.New(CacheSize)

	if (err != nil) { panic(err) }

	client := new(StatCounterClient)
	client.HttpClient = new(http.Client);
	client.BaseUrl = "http://gs.statcounter.com/chart.php?bar=1&device=Desktop%2C%20Tablet%20%26%20Console&device_hidden=desktop%2Btablet%2Bconsole&multi-device=true&statType_hidden=browser_version&region_hidden=ww&granularity=monthly&statType=Browser%20Version&region=Worldwide"
	client.Cache = cache

	return client
}

func (client *StatCounterClient) GetBrowserShare(fromDate string, toDate string, channel chan *StatCounterResponse) {
	url := fmt.Sprintf("%v&fromMonthYear=%v&toMonthYear=%v", client.BaseUrl, fromDate, toDate)
	cachedResponse, _ := client.Cache.Get(url)
	
	if cachedResponse != nil { channel <- NewStatCounterResponse(cachedResponse.(*StatCounterChart), nil) }
	
	response, err := client.HttpClient.Get(url)

	if err != nil { channel <- NewStatCounterResponse(nil, err) }

	statCounterChart, deserialiseErr := client.Deserialise(response.Body)
	client.Cache.Add(url, statCounterChart)
	channel <- NewStatCounterResponse(statCounterChart, deserialiseErr)
}

func (client *StatCounterClient) Deserialise(body io.ReadCloser) (*StatCounterChart, error) {
	var statCounterChart *StatCounterChart

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(body)
	err := xml.Unmarshal(buffer.Bytes(), &statCounterChart)
	body.Close()

	return statCounterChart, err
}