package tests

import (
	"testing"
	"sort"
	"browserstats/internal"
)

/* Go ranges don't honour ordering, so this 
 * is a workaround to keep the array
 * order consistent during tests.
 * Not sorting in production to save on
 * CPU overhead */

type SortableBrowsers []*internal.Browser

func (b SortableBrowsers) Len() (int) {
	return len(b)
}

func (b SortableBrowsers) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b SortableBrowsers) Less(i, j int) bool {
	return b[i].Name > b[j].Name
}

func TestResponseFactory(t *testing.T) {
	statCounterChart := new(internal.StatCounterChart)
	firefox := new(internal.StatCounterSet)
	chrome := new(internal.StatCounterSet)

	firefox.Label = "Firefox"
	firefox.Value = "2.25"

	chrome.Label = "Chrome"
	chrome.Value = "49"

	statCounterChart.Browsers = []*internal.StatCounterSet {
		chrome,
		firefox,
	}

	browsers := SortableBrowsers(internal.NewResponse(statCounterChart))

	sort.Sort(browsers);
	
	if len(browsers) == 0 { t.Error("No browsers present in aggregated array") }

	if browsers[0].Name != "Firefox" {
		t.Error("Expected fist name to be \"Firefox\", not", browsers[0].Name)
	}

	if browsers[0].Share != "2.25" {
		t.Error("Expected first share to be \"2.25\", not", browsers[0].Share)
	}

	if browsers[1].Name != "Chrome" {
		t.Error("Expected second name to be \"Chrome\", not", browsers[1].Name)
	}

	if browsers[1].Share != "49" {
		t.Error("Expected second share to be \"49\", not", browsers[1].Share)
	}
}