package models

import "strconv"

type IBaseError interface {
	Code() int
	Message() string
	Error() string
}

type BaseError struct {
	code    int
	message string
}

func newBaseError(code int, message string) (baseError *BaseError) {
	baseError = new(BaseError)
	baseError.code = code
	baseError.message = message
	return
}

func (e *BaseError) Code() int {
	return e.code
}

func (e *BaseError) Message() string {
	return e.message
}

func (e *BaseError) Error() string {
	return strconv.Itoa(e.code) + " " + e.message
}

// 400~

type ForbiddenError struct {
	*BaseError
}

func NewForbiddenError(message string) (error *ForbiddenError) {
	error = new(ForbiddenError)
	error.BaseError = newBaseError(403, "ForbiddenError:"+message)
	return
}

type NotFoundError struct {
	*BaseError
}

func NewNotFoundError(message string) (error *NotFoundError) {
	error = new(NotFoundError)
	error.BaseError = newBaseError(404, "Not Found Error:"+message)
	return
}

// 500~

type InternalServerError struct {
	*BaseError
}

func NewInternalServerError(message string) (error *InternalServerError) {
	error = new(InternalServerError)
	error.BaseError = newBaseError(500, "Internal Server Error:"+message)
	return
}
