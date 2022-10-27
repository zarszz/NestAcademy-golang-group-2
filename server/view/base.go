package view

import "net/http"

type Query struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Total int `json:"total"`
}

type ResponseSuccess struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	GeneralInfo string `json:"general_info"`
}

type ResponseWithDataSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

type ResponseGetPaginationSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
	Query   Query       `json:"query"`
}

type ResponseFailed struct {
	Status         int         `json:"status"`
	Message        string      `json:"message"`
	Error          string      `json:"error"`
	AdditionalInfo interface{} `json:"additional_info"`
	GeneralInfo    string      `json:"general_info"`
}

type AdditionalInfoError struct {
	Message interface{} `json:"message"`
}

func SuccessCreated(message string) *ResponseSuccess {
	return &ResponseSuccess{
		Status:      http.StatusCreated,
		Message:     message,
		GeneralInfo: "Group-2",
	}
}

func SuccessWithData(payload interface{}, message string) *ResponseWithDataSuccess {
	return &ResponseWithDataSuccess{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func SuccessGetPagination(payload interface{}, message string, query Query) *ResponseGetPaginationSuccess {
	return &ResponseGetPaginationSuccess{
		Status:  http.StatusOK,
		Payload: payload,
		Message: message,
		Query:   query,
	}
}

func ErrBadRequest(additionalInfo interface{}, message string) *ResponseFailed {
	return &ResponseFailed{
		Status:         http.StatusBadRequest,
		AdditionalInfo: additionalInfo,
		Message:        message,
		Error:          "BAD_REQUEST",
		GeneralInfo:    "Kelompok-2",
	}
}

func ErrInternalServer(additionalInfo interface{}, message string) *ResponseFailed {
	return &ResponseFailed{
		Status:         http.StatusInternalServerError,
		AdditionalInfo: additionalInfo,
		Message:        message,
		Error:          "INTERNAL_SERVER_ERROR",
		GeneralInfo:    "Kelompok-2",
	}
}

func ErrNotFound(additionalInfo interface{}, message string) *ResponseFailed {
	return &ResponseFailed{
		Status:         http.StatusNotFound,
		AdditionalInfo: additionalInfo,
		Message:        message,
		Error:          "NOT_FOUND",
		GeneralInfo:    "Kelompok-2",
	}
}

func ErrUnauthorized(additionalInfo interface{}, message string) *ResponseFailed {
	return &ResponseFailed{
		Status:         http.StatusUnauthorized,
		AdditionalInfo: additionalInfo,
		Message:        message,
		Error:          "UNAUTHORIZED",
		GeneralInfo:    "Kelompok-2",
	}
}

func InvalidRequestPayload() *AdditionalInfoError {
	return &AdditionalInfoError{
		Message: "invalid request payload",
	}
}