package decoder

type CsvDecoder struct {
	validator Validator
}

func NewCsvDecoder(validator Validator) *CsvDecoder {
	return &CsvDecoder{
		validator: validator,
	}
}

func (p *CsvDecoder) Decode(input interface{}) ([]Record, []Record) {
	var records []Record
	var failedRecords []Record

	// We assume the CSV input is being parsed into a matrix
	if csv, ok := input.([][]string); ok {
		for _, row := range csv {
			record := p.validator.Validate(row)
			if record.Error != nil {
				failedRecords = append(failedRecords, record)
			} else {
				records = append(records, record)
			}

		}
	}

	return records, failedRecords
}
