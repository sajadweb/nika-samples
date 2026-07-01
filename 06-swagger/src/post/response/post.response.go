package response

import "nikaapp/src/post/services"

 
 type ErrorDetail struct {
	Code    int      `json:"code"`              // error code (e.g.: VALIDATION_ERROR, USER_NOT_FOUND)
	Message string      `json:"message"`           // error message
	Details interface{} `json:"details,omitempty"` // details (populated only for validation errors)
}
type Error struct {
	Success bool         `json:"success"`
	Message string       `json:"message,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

type CreateResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    services.Post `json:"data"`
}
type DeleteResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"` 
}

type ListResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []services.Post `json:"data"`
	Meta    Meta           `json:"meta"`
}

type Meta struct {
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Count int64 `json:"count"`
}
