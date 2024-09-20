package processor

import (
	loader2 "github.com/juaninterviews/stori-tech-interview/pkg/loader"
	"github.com/juaninterviews/stori-tech-interview/pkg/worker"
	"github.com/juaninterviews/stori-tech-interview/pkg/worker/decoder"
	"io"
	"sync"
)

type Processor interface {
	ProcessFileChunked(r io.Reader) *FileResult
}

type FileProcessor struct {
	loader loader2.Strategy
	worker worker.IWorker
}

func NewFileLoader(
	loader loader2.Strategy,
	worker worker.IWorker,
) *FileProcessor {
	return &FileProcessor{
		loader: loader,
		worker: worker,
	}
}

type FileResult struct {
	FailedRecords       []decoder.Record
	SuccessRecordsCount int
	InternalErrors      []error
	InternalError       error
}

func (f *FileProcessor) ProcessFileChunked(r io.Reader) *FileResult {
	records, err := f.loader.LoadFileChunked(r)
	if err != nil {
		return &FileResult{
			FailedRecords:       nil,
			SuccessRecordsCount: 0,
			InternalErrors:      nil,
			InternalError:       err,
		}
	}

	var wg sync.WaitGroup
	var workersResults []*worker.Result
	for e := records.Front(); e != nil; e = e.Next() {
		wg.Add(1)

		e := e
		go func() {
			defer wg.Done()
			workerRes := f.worker.Process(e.Value)
			workersResults = append(workersResults, workerRes)
		}()
	}
	wg.Wait()

	result := &FileResult{
		FailedRecords:       []decoder.Record{},
		SuccessRecordsCount: 0,
		InternalError:       nil,
		InternalErrors:      []error{},
	}
	for _, workerResult := range workersResults {
		result.FailedRecords = append(result.FailedRecords, workerResult.FailedRecords...)
		result.SuccessRecordsCount += workerResult.SuccessRecordsCount
		result.InternalErrors = append(result.InternalErrors, workerResult.InternalError)
	}

	return result
}
