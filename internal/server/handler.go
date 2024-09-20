package server

import (
	"encoding/json"
	"github.com/juaninterviews/stori-tech-interview/internal/uploader"
	"github.com/juaninterviews/stori-tech-interview/pkg/processor"
	"net/http"
)

type IHandler interface {
	MigrateHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	UploadManager uploader.IUploaderManager
}

func NewHandler() *Handler {
	return &Handler{
		UploadManager: uploader.NewUploadManager(),
	}
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

type ErrorRecord struct {
	FailedRecords []struct {
		Record map[string]interface{} `json:"record"`
		Error  string                 `json:"error"`
	} `json:"failedRecords"`
	SuccessRecordsCount int `json:"successRecordsCount"`
}

func (h *Handler) MigrateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		httpError(w, http.StatusBadRequest, "Error file processing", err.Error())
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		httpError(w, http.StatusBadRequest, "File not exist with key 'file'", err.Error())
		return
	}
	defer file.Close()

	fileError := h.UploadManager.MigrateCSVFile(file)
	if err != nil {
		httpError(w, http.StatusBadRequest, "Internal error", err.Error())
	}

	failedRecordsHandler(w, fileError)
}

func httpError(w http.ResponseWriter, statusCode int, message string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := ErrorResponse{Error: message, Details: details}
	json.NewEncoder(w).Encode(errorResponse)
}

func failedRecordsHandler(w http.ResponseWriter, result *processor.FileResult) {
	var failedRecords []struct {
		Record map[string]interface{} `json:"record"`
		Error  string                 `json:"error"`
	}

	for _, record := range result.FailedRecords {
		processedRecord := struct {
			Record map[string]interface{} `json:"record"`
			Error  string                 `json:"error"`
		}{
			Record: record.Record,
			Error:  record.Error.Error(),
		}
		failedRecords = append(failedRecords, processedRecord)
	}

	response := ErrorRecord{
		FailedRecords:       failedRecords,
		SuccessRecordsCount: result.SuccessRecordsCount,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
