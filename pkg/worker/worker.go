package worker

import "github.com/juaninterviews/stori-tech-interview/pkg/worker/decoder"

type IWorker interface {
	Process(input interface{}) *Result
}

type Worker struct {
	decoder decoder.Decoder
	action  func([]decoder.Record) error
}

type Result struct {
	FailedRecords       []decoder.Record
	SuccessRecordsCount int
	InternalError       error
}

func NewWorker(
	decoder decoder.Decoder,
	action func([]decoder.Record) error,
) *Worker {
	return &Worker{
		decoder: decoder,
		action:  action,
	}
}

func (w *Worker) Process(input interface{}) *Result {
	var result = &Result{
		FailedRecords:       nil,
		SuccessRecordsCount: 0,
		InternalError:       nil,
	}

	records, failed := w.decoder.Decode(input)

	err := w.action(records)
	if err != nil {
		result.InternalError = err
		result.FailedRecords = append(records, failed...)
		result.SuccessRecordsCount = 0
		return result
	}

	result.SuccessRecordsCount = len(records)
	result.FailedRecords = failed
	return result
}
