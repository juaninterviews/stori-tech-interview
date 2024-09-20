package loader

import (
	"container/list"
	"encoding/csv"
	"io"
)

type Strategy interface {
	LoadFileChunked(r io.Reader) (*list.List, error)
}

type Csv struct {
	PageSize int
}

func NewCsv(pageSize int) *Csv {
	return &Csv{
		PageSize: pageSize,
	}
}

func (c *Csv) LoadFileChunked(r io.Reader) (*list.List, error) {
	reader := csv.NewReader(r)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	min := 1
	max := c.PageSize + 1
	if max > len(records) {
		max = len(records)
	}

	pages := (len(records) - 1 + c.PageSize - 1) / c.PageSize
	if pages < 1 {
		pages = 1
	}

	var chunks list.List
	for i := 0; i < pages; i++ {
		chunk := records[min:max]
		chunks.PushBack(chunk)

		min = max
		max = max + c.PageSize
		if max > len(records) {
			max = len(records)
		}
	}

	return &chunks, nil
}
