package loader

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type CsvLoaderSuite struct {
	suite.Suite
}

func TestCsvLoaderSuite(t *testing.T) {
	suite.Run(t, new(CsvLoaderSuite))
}

func (suite *CsvLoaderSuite) TestShouldLoadPaginatedDocuments() {
	loader := NewCsv(10)

	in := `id,user_id,amount,datetime
1,2,10.00,2024-07-01T15:00:12Z
2,1,-5.00,2024-07-01T15:00:12Z 
3,4,-5.52,2024-06-01T14:59:59Z 
4,1,21.27,2024-07-01T12:00:12Z 
5,1,15.00,2024-06-01T13:59:59Z`

	chunked, err := loader.LoadFileChunked(strings.NewReader(in))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), chunked.Len(), 1)
}

func (suite *CsvLoaderSuite) TestShouldLoadPaginatedDocumentsOnePageSize() {
	loader := NewCsv(1)

	in := `id,user_id,amount,datetime
1,2,10.00,2024-07-01T15:00:12Z
2,1,-5.00,2024-07-01T15:00:12Z 
3,4,-5.52,2024-06-01T14:59:59Z 
4,1,21.27,2024-07-01T12:00:12Z 
5,1,15.00,2024-06-01T13:59:59Z`

	chunked, err := loader.LoadFileChunked(strings.NewReader(in))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), chunked.Len(), 5)
}

func (suite *CsvLoaderSuite) TestShouldLoadPaginatedDocumentsOnTwoPages() {
	loader := NewCsv(5)

	in := `id,user_id,amount,datetime
1,2,10.00,2024-07-01T15:00:12Z
2,1,-5.00,2024-07-01T15:00:12Z 
3,4,-5.52,2024-06-01T14:59:59Z 
4,1,21.27,2024-07-01T12:00:12Z 
5,1,15.00,2024-06-01T13:59:59Z
6,2,10.00,2024-07-01T15:00:12Z
7,1,-5.00,2024-07-01T15:00:12Z 
8,4,-5.52,2024-06-01T14:59:59Z 
9,1,21.27,2024-07-01T12:00:12Z 
10,1,15.00,2024-06-01T13:59:59Z`

	chunked, err := loader.LoadFileChunked(strings.NewReader(in))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), chunked.Len(), 2)
}

func (suite *CsvLoaderSuite) TestShouldLoadPaginatedDocumentsOnThreePagesWithRemain() {
	loader := NewCsv(5)

	in := `id,user_id,amount,datetime
1,2,10.00,2024-07-01T15:00:12Z
2,1,-5.00,2024-07-01T15:00:12Z 
3,4,-5.52,2024-06-01T14:59:59Z 
4,1,21.27,2024-07-01T12:00:12Z 
5,1,15.00,2024-06-01T13:59:59Z
6,2,10.00,2024-07-01T15:00:12Z
7,1,-5.00,2024-07-01T15:00:12Z 
8,4,-5.52,2024-06-01T14:59:59Z 
9,1,21.27,2024-07-01T12:00:12Z 
10,1,15.00,2024-06-01T13:59:59Z
11,1,15.00,2024-06-01T13:59:59Z`

	chunked, err := loader.LoadFileChunked(strings.NewReader(in))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), chunked.Len(), 3)
}
