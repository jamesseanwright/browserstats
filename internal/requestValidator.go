package internal

import (
	"errors"
	"regexp"
)

const FromParamKey = "from"
const ToParamKey = "to"

var dateFormat = regexp.MustCompile(`[0-9]{4}-[0-9]{2}`)

type RequestValidator struct {}

func NewRequestValidator() (*RequestValidator) {
	return new(RequestValidator)
}

func (requestValidator *RequestValidator) Validate(query map[string][]string) (string, string, error) {
	fromParam := query[FromParamKey]
	toParam := query[ToParamKey]

	if fromParam == nil {
		return "", "", errors.New("from parameter is missing")
	}

	if toParam == nil {
		return "", "", errors.New("to parameter is missing")
	}

	if len(fromParam) > 1 || !dateFormat.MatchString(fromParam[0]) {
		return "", "", errors.New("from date format is invalid")
	}

	if len(toParam) > 1 || !dateFormat.MatchString(toParam[0]) {
		return "", "", errors.New("to date format is invalid")
	}

	return fromParam[0], toParam[0], nil
}