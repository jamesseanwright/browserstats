package internal

import (
	"fmt"
)

type ResponseBuilder struct {}

func NewResponseBuilder() (*ResponseBuilder){
	return new(ResponseBuilder)
}

func (responseBuilder *ResponseBuilder) Build(statCounterChart *StatCounterChart) {
	fmt.Println(statCounterChart)
}