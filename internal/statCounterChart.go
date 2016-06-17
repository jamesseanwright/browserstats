package internal

type StatCounterChart struct {
	Browsers []StatCounterSet `xml:"chart>set"`
}