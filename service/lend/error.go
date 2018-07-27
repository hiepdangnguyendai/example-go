package lend

import (
	"net/http"
)

var (
	ErrNotFound         = errNotFound{}
	ErrUnknown          = errUnknown{}
	ErrUserIDIsRequired = errUserIDIsRequired{}
	ErrUserIsNotExisted = errUserIsNotExisted{}
	ErrBookIDIsRequired = errBookIDIsRequired{}
	ErrBookIsNotExisted = errBookIsNotExisted{}
	ErrToIsRequired     = errToIsRequired{}
	ErrRecordNotFound   = errRecordNotFound{}
	ErrBookNotAvailable = errBookNotAvailable{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUserIDIsRequired struct{}

func (errUserIDIsRequired) Error() string {
	return "UserID is required"
}

func (errUserIDIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errUserIsNotExisted struct{}

func (errUserIsNotExisted) Error() string {
	return "User is not existed"
}

func (errUserIsNotExisted) StatusCode() int {
	return http.StatusBadRequest
}

type errBookIDIsRequired struct{}

func (errBookIDIsRequired) Error() string {
	return "BookID is required"
}

func (errBookIDIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errBookIsNotExisted struct{}

func (errBookIsNotExisted) Error() string {
	return "Book is not existed"
}

func (errBookIsNotExisted) StatusCode() int {
	return http.StatusBadRequest
}

type errToIsRequired struct{}

func (errToIsRequired) Error() string {
	return "required"
}

func (errToIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errBookNotAvailable struct{}

func (errBookNotAvailable) Error() string {
	return " book not available"
}

func (errBookNotAvailable) StatusCode() int {
	return http.StatusBadRequest
}
