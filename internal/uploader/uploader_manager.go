package uploader

import (
	"fmt"
	"github.com/juaninterviews/stori-tech-interview/pkg/balance"
	loader2 "github.com/juaninterviews/stori-tech-interview/pkg/loader"
	"github.com/juaninterviews/stori-tech-interview/pkg/processor"
	"github.com/juaninterviews/stori-tech-interview/pkg/worker"
	"github.com/juaninterviews/stori-tech-interview/pkg/worker/decoder"
	"io"
)

type IUploaderManager interface {
	MigrateCSVFile(r io.Reader) *processor.FileResult
}

type UploadManager struct {
	BalanceManager IBalanceManager
}

func NewUploadManager() *UploadManager {
	return &UploadManager{
		BalanceManager: NewBalanceManager(),
	}
}

func (m *UploadManager) MigrateCSVFile(r io.Reader) *processor.FileResult {
	loader := processor.NewFileLoader(
		loader2.NewCsv(50),
		worker.NewWorker(
			decoder.NewCsvDecoder(
				balance.NewValidator(),
			),
			func(input []decoder.Record) error {
				//m.BalanceManager.InsertBalance()
				for _, record := range input {
					fmt.Printf("Input %v, error %v\n", record.Record, record.Error)
				}
				return nil
			},
		),
	)

	failed := loader.ProcessFileChunked(r)

	return failed
}
