package decoder

type Record struct {
	Record map[string]interface{} `json:"record"`
	Error  error                  `json:"error"`
}

type Decoder interface {
	Decode(input interface{}) ([]Record, []Record)
}

type Validator interface {
	Validate(input interface{}) Record
}
