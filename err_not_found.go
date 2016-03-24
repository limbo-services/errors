package errors

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// NotFoundError should be used when some infomation could not be found
type NotFoundError interface {
	error
	IsNotFound() bool
}

// IsNotFound returns true when:
// - err is a NotFoundError
// - os.IsNotExist is true
// - err equals sql.ErrNoRows
// - grpc.Code is codes.NotFound
func IsNotFound(err error) bool {
	if err != nil {
		if err == sql.ErrNoRows {
			return true
		}
		if os.IsNotExist(err) {
			return true
		}
		if grpc.Code(err) == codes.NotFound {
			return true
		}
		x, _ := err.(NotFoundError)
		if x != nil && x.IsNotFound() {
			return true
		}
	}
	return false
}

// NotFound makes a new NotFoundError
func NotFound(a ...interface{}) error {
	return &errNotFound{s: fmt.Sprint(a...)}
}

// NotFoundf is the formatted version of NotFound
func NotFoundf(format string, a ...interface{}) error {
	return &errNotFound{s: fmt.Sprintf(format, a...)}
}

type errNotFound struct{ s string }

func (e *errNotFound) Error() string        { return e.s }
func (e *errNotFound) IsNotFound() bool     { return true }
func (e *errNotFound) HTTPStatusCode() int  { return http.StatusNotFound }
func (e *errNotFound) HTTPMessage() string  { return e.Error() }
func (e *errNotFound) GRPCCode() codes.Code { return codes.NotFound }
func (e *errNotFound) GRPCMessage() string  { return e.Error() }
