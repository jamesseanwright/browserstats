package internal

type Browser struct {
	Name string `json:"name"`
	Share string `json:"share"`
}


func NewResponse(statCounterChart *StatCounterChart) ([]*Browser) {
	browsers := make([]*Browser, len(statCounterChart.Browsers))

	for i, statCounterSet := range statCounterChart.Browsers {
		browser := new(Browser)
		browser.Name = statCounterSet.Label
		browser.Share = statCounterSet.Value

		browsers[i] = browser
	}

	return browsers
}