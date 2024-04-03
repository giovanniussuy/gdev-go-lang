package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/giovanniussuy/gdev-go-lang/app/audit"
)

type MockHTTPClient struct {
	Response *http.Response
	Error    error
}

type Request[TRequest any] struct {
	Body    *TRequest
	TraceId string
	Context context.Context
}

type Response[TResponse any] struct {
	AuditResponse    audit.IAuditResponseStatus
	Data             *TResponse
	OriginalResponse *http.Response
}

type APIError string

const (
	ServiceError = "SE1000"
)

func getData(url string, traceId string, mockClient *MockHTTPClient) (*http.Response, error) {
	if mockClient != nil {
		return mockClient.Response, mockClient.Error
	}

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("traceId", traceId)

	return http.DefaultClient.Do(request)
}

func Post[TRequest any, TResponse any](path string, request Request[TRequest]) *Response[TResponse] {
	bodyBytes, err := json.Marshal(request.Body)

	// need to change log to more dynamic func
	if err != nil {
		log.Printf("ERROR: Payload bytes conversion error: %+v, error: %s", request.Body, err)
		return nil

	}

	payload := bytes.NewBuffer(bodyBytes)
	httpRequest, _ := http.NewRequest("POST", path, payload)
	httpRequest.Header.Add("traceId", request.TraceId)

	return handleResponse[TRequest, TResponse](httpRequest)
}

// request and handle the response
func handleResponse[TRequest any, TResponse any](httpRequest *http.Request) *Response[TResponse] {
	response, err := http.DefaultClient.Do(httpRequest)

	if err != nil {
		log.Printf("Request to %s with response %+v, error: %+v", httpRequest.URL, response, err)

		return &Response[TResponse]{
			AuditResponse:    audit.ConstructAuditResponseStatusByCode(audit.ServiceError),
			Data:             nil,
			OriginalResponse: response,
		}
	}

	log.Printf("%s request: %+v", httpRequest.URL, response)
	actualResponse, mapError, err := ConvertReaderToType[TResponse](response.Body, *new(TResponse))

	if mapError != nil {
		log.Printf("Response model %+v conversion with error: %s", response.Body, err)

		return &Response[TResponse]{
			AuditResponse:    audit.ConstructAuditResponseStatusByCode(audit.ServiceError),
			Data:             nil,
			OriginalResponse: response,
		}

	}

	if response.StatusCode >= 300 {
		log.Printf("Unespered status code on %s error: %+v", httpRequest.URL, response)

		return &Response[TResponse]{
			AuditResponse:    audit.ConstructAuditResponseStatusByCode(audit.ServiceError),
			Data:             nil,
			OriginalResponse: response,
		}
	}

	log.Printf("Request with success %s response: %+v", httpRequest.URL, actualResponse)
	return &Response[TResponse]{
		AuditResponse:    nil,
		Data:             &actualResponse,
		OriginalResponse: response,
	}

}

func ConvertReaderToType[T any](requestResponse io.Reader, genericType T) (T, audit.IAuditResponseStatus, error) {
	byteBody, readAllError := io.ReadAll(requestResponse)

	if readAllError != nil {
		return *new(T), audit.ConstructAuditResponseStatusByCode(audit.ServiceError), readAllError
	}

	if unmarshalError := json.Unmarshal(byteBody, &genericType); unmarshalError != nil {
		return *new(T), audit.ConstructAuditResponseStatusByCode(audit.ServiceError), unmarshalError
	}

	return genericType, nil, nil
}
