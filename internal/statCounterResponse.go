package internal

type StatCounterResponse struct {
	Response *StatCounterChart
	Error error
}

func NewStatCounterResponse(response *StatCounterChart, err error) (*StatCounterResponse) {
	statCounterResponse := new(StatCounterResponse)
	statCounterResponse.Response = response
	statCounterResponse.Error = err

	return statCounterResponse
}