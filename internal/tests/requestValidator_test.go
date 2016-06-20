package tests

import (
	"testing"
	"browserstats/internal"
)

func TestValidateMissingFrom(t *testing.T) {
	validator := internal.NewRequestValidator()
	query := make(map[string][]string)

	from, to, err := validator.Validate(query)

	if from != "" { t.Error("Expected returned from to be empty") }
	if to != "" { t.Error("Expected returned to to be empty") }
	if err == nil { t.Error("Expected err to not be nil") }

	if err.Error() != "from parameter is missing" { 
		t.Error("Expected mesage to be \"from parameter is missing\", not", err.Error())
	}
}

func TestValidateMissingTo(t *testing.T) {
	validator := internal.NewRequestValidator()

	query := map[string][]string{
		"from": []string{ "2016-06" },
	}

	from, to, err := validator.Validate(query)

	if from != "" { t.Error("Expected returned from to be empty") }
	if to != "" { t.Error("Expected returned to to be empty") }
	if err == nil { t.Error("Expected err to not be nil") }

	if err.Error() != "to parameter is missing" { 
		t.Error("Expected mesage to be \"to parameter is missing\", not", err.Error())
	}
}

func TestValidateMultipleFromParams(t *testing.T) {
	validator := internal.NewRequestValidator()

	query := map[string][]string{
		"from": []string{ "2016-06", "2016-07" },
		"to": []string{ "2016-06" },		
	}

	from, to, err := validator.Validate(query)

	if from != "" { t.Error("Expected returned from to be empty") }
	if to != "" { t.Error("Expected returned to to be empty") }
	if err == nil { t.Error("Expected err to not be nil") }

	if err.Error() != "from date format is invalid" { 
		t.Error("Expected mesage to be \"from date format is invalid\", not", err.Error())
	}
}

func TestValidateMultipleToParams(t *testing.T) {
	validator := internal.NewRequestValidator()

	query := map[string][]string{
		"from": []string{ "2016-06" },
		"to": []string{ "2016-06", "2016-07" },
	}

	from, to, err := validator.Validate(query)

	if from != "" { t.Error("Expected returned from to be empty") }
	if to != "" { t.Error("Expected returned to to be empty") }
	if err == nil { t.Error("Expected err to not be nil") }

	if err.Error() != "to date format is invalid" { 
		t.Error("Expected mesage to be \"to date format is invalid\", not", err.Error())
	}
}

func TestValidateInvalidFromDataFormat(t *testing.T) {
	validator := internal.NewRequestValidator()

	query := map[string][]string{
		"from": []string{ "foobarbaz" },
		"to": []string{ "2016-06" },		
	}

	from, to, err := validator.Validate(query)

	if from != "" { t.Error("Expected returned from to be empty") }
	if to != "" { t.Error("Expected returned to to be empty") }
	if err == nil { t.Error("Expected err to not be nil") }

	if err.Error() != "from date format is invalid" { 
		t.Error("Expected mesage to be \"from date format is invalid\", not", err.Error())
	}
}

func TestValidateInvalidToDataFormat(t *testing.T) {
	validator := internal.NewRequestValidator()

	query := map[string][]string{
		"from": []string{ "2016-06" },	
		"to": []string{ "foobarbaz" },			
	}

	from, to, err := validator.Validate(query)

	if from != "" { t.Error("Expected returned from to be empty") }
	if to != "" { t.Error("Expected returned to to be empty") }
	if err == nil { t.Error("Expected err to not be nil") }

	if err.Error() != "to date format is invalid" { 
		t.Error("Expected mesage to be \"to date format is invalid\", not", err.Error())
	}
}

func TestValidateCorrectQuery(t *testing.T) {
	validator := internal.NewRequestValidator()

	query := map[string][]string{
		"from": []string{ "2016-06" },	
		"to": []string{ "2016-07" },			
	}

	from, to, err := validator.Validate(query)

	if from == "" { t.Error("Expected returned from to not be empty") }
	if to == "" { t.Error("Expected returned to to not be empty") }
	if err != nil { t.Error("Expected err to be nil") }

	if from != "2016-06" { 
		t.Error("Expected from to be \"2016-06\", not", from)
	}

	if to != "2016-07" { 
		t.Error("Expected to to be \"2016-07\", not", to)
	}
}