package balance

import (
	"fmt"
	"github.com/juaninterviews/stori-tech-interview/pkg/worker/decoder"
	"strconv"
	"strings"
	"time"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (b *Validator) Validate(input interface{}) decoder.Record {
	row, ok := input.([]string)
	if !ok {
		return decoder.Record{
			Record: nil,
			Error:  fmt.Errorf("input has wrong format"),
		}
	}

	id := strings.Trim(row[0], " ")
	userID := strings.Trim(row[1], " ")
	amount := strings.Trim(row[2], " ")
	datetime := strings.Trim(row[3], " ")

	data := decoder.Record{
		Record: map[string]interface{}{
			"id":        id,
			"userId":    userID,
			"amount":    amount,
			"timestamp": datetime,
		},
		Error: nil,
	}

	var errorsList []error
	if id == "" {
		errorsList = append(errorsList, fmt.Errorf("id is null or invalid"))
	}
	if userID == "" {
		errorsList = append(errorsList, fmt.Errorf("userID is null or invalid"))
	}

	//	Amount must be floating number
	if amount == "" {
		errorsList = append(errorsList, fmt.Errorf("amount is null or invalid"))
	} else {
		if _, err := strconv.ParseFloat(amount, 64); err != nil {
			errorsList = append(errorsList, fmt.Errorf("amount is not a valid number"))
		}
	}

	// Datetime validation (ISO 8601)
	if datetime == "" {
		errorsList = append(errorsList, fmt.Errorf("datetime is null or invalid"))
	} else {
		_, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			errorsList = append(errorsList, fmt.Errorf("datetime format is invalid, expected ISO 8601"))
		}
	}

	if len(errorsList) > 0 {
		data.Error = fmt.Errorf("multiple errors: %v", errorsList)
		return data
	}

	return data
}
