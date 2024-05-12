package audit

import (
	"fmt"
	"net/http"
)

type IAuditResponseStatus interface {
	GetCode() string
	GetMessage() string
	GetStatus() int
	SetCode(code string)
	SetMessage(message string)
	SetStatus(status int)
}

func (a *AuditResponseStatus) GetCode() string {
	return a.Code
}

func (a *AuditResponseStatus) GetMessage() string {
	return a.Message
}
func (a *AuditResponseStatus) GetStatus() int {
	return a.Status
}
func (a *AuditResponseStatus) SetCode(code string) {
	a.Code = code
}
func (a *AuditResponseStatus) SetMessage(message string) {
	a.Message = message
}
func (a *AuditResponseStatus) SetStatus(status int) {
	a.Status = status
}

type AuditResponseStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
	Error   error  `json:"-"`
}

type AuditCode string

var AuditMEssages = map[AuditCode]struct {
	Code    string
	Message string
	Status  int
}{
	STATUS_OK: {STATUS_OK, "success", http.StatusOK},

	BAD_REQUEST: {BAD_REQUEST, "error", http.StatusBadRequest},

	INTERNAL_ERROR: {INTERNAL_ERROR, "error", http.StatusInternalServerError},
}

const (
	STATUS_OK      = "CS0001"
	BAD_REQUEST    = "SE1000"
	INTERNAL_ERROR = "IE1000"
)

func ConstructAuditResponseStatusByCode(code AuditCode) IAuditResponseStatus {
	audit, find := AuditMEssages[code]

	if !find {
		return &AuditResponseStatus{
			Code:    INTERNAL_ERROR,
			Message: "error.",
			Status:  http.StatusInternalServerError,
		}
	}

	return &AuditResponseStatus{
		Code:    audit.Code,
		Message: audit.Message,
		Status:  audit.Status,
	}
}

func ConstructAuditResponseStatusByCodeWithParams(code AuditCode, extraMessage string) IAuditResponseStatus {
	auditResponse := ConstructAuditResponseStatusByCode(code)

	auditResponse.SetMessage((fmt.Sprintf("%s: %s.", auditResponse.GetMessage(), extraMessage)))

	return auditResponse
}

func ConstructAuditResponse(code string, message string, status int, err error) IAuditResponseStatus {
	return &AuditResponseStatus{
		Code:    code,
		Message: message,
		Status:  status,
		Error:   err,
	}
}
