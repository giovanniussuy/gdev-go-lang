package client

import (
	"encoding/json"

	"github.com/giovanniussuy/gdev-go-lang/app/audit"
	"github.com/giovanniussuy/gdev-go-lang/app/client/fasthttp"
	"github.com/giovanniussuy/gdev-go-lang/app/model"
)

/*
const (
	xAmaznTraceIdField = "x-amzn-trace-id"
	xApigwApiId        = "x-apigw-api-id"
	xApiKey            = "x-api-key"
	appJson            = "application/json"
)*/

func HttpSend[T any](clientRequest fasthttp.Request, traceId string, responseType T) (T, audit.IAuditResponseStatus) {
	response, requetError := fasthttp.Init().Send(&clientRequest)

	if requetError != nil {
		return *new(T), audit.ConstructAuditResponseStatusByCode(audit.INTERNAL_ERROR)
	}

	if response.StatusCode >= 400 {
		return *new(T), convertErrorResponse(response)
	}

	convertedResponse, auditError, _ := ConvertReaderToType(response.Body, responseType)

	if auditError != nil {
		return *new(T), auditError
	}

	return convertedResponse, nil
}

func convertErrorResponse(response fasthttp.Response) audit.IAuditResponseStatus {
	errorResponse, auditError, _ := ConvertReaderToType(response.Body, model.ErrorResponse{})
	errorResponse.Status = response.StatusCode

	if errorResponse.Descricao == "" {
		return audit.ConstructAuditResponseStatusByCode(audit.INTERNAL_ERROR)
	}

	if auditError != nil {
		return auditError
	}

	return audit.ConstructAuditResponse(errorResponse.Codigo, errorResponse.Descricao, auditError.GetStatus(), nil)
}

func ConvertReaderToType[T any](byteBody []byte, genericType T) (T, audit.IAuditResponseStatus, error) {
	if unmarshalError := json.Unmarshal(byteBody, &genericType); unmarshalError != nil {
		return *new(T), audit.ConstructAuditResponseStatusByCode(audit.INTERNAL_ERROR), unmarshalError
	}
	return genericType, nil, nil
}
