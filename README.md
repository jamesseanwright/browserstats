# BrowserStats

A HTTP API written in Go that exposes StatsCounter's browser share dataset as JSON 


## Production Host

http://browser-stats.herokuapp.com/stats?from=2016-06&to=2016-06


## API

### `GET /stats`

Returns an array of the most popular browsers and respective market shares.

#### Query Parameters

* `from` - the month and year at which the data should start
* `to` - the month and year at which the data should end

These parameters must be specified in a `YYYY-MM` format, otherwise one will receive a HTTP 400.


## Local Development

* `go get github.com/tools/godep`
* `cd <your Go src directory>`
* `git clone https://github.com/jamesseanwright/browserstats.git`
* `cd browserstats`
* `godep restore`
* `PORT=8080 go run server.go`


## Unit Tests

They're on their way!
