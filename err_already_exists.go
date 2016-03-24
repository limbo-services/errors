package errors

import (
	"fmt"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// AlreadyExistsError should be used when some infomation should not be found
type AlreadyExistsError interface {
	error
	IsAlreadyExists() bool
}

// IsAlreadyExists returns true when:
// - err is a AlreadyExistsError
// - os.IsExist is true
// - if grpc.Code() is codes.AlreadyExists
func IsAlreadyExists(err error) bool {
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if grpc.Code(err) == codes.AlreadyExists {
			return true
		}
		x, _ := err.(AlreadyExistsError)
		if x != nil && x.IsAlreadyExists() {
			return true
		}
	}
	return false
}

// AlreadyExists makes a new AlreadyExistsError
func AlreadyExists(a ...interface{}) error {
	return &errAlreadyExists{s: fmt.Sprint(a...)}
}

// AlreadyExistsf is the formatted version of AlreadyExists
func AlreadyExistsf(format string, a ...interface{}) error {
	return &errAlreadyExists{s: fmt.Sprintf(format, a...)}
}

type errAlreadyExists struct{ s string }

func (e *errAlreadyExists) Error() string         { return e.s }
func (e *errAlreadyExists) IsAlreadyExists() bool { return true }
func (e *errAlreadyExists) HTTPStatusCode() int   { return http.StatusConflict }
func (e *errAlreadyExists) HTTPMessage() string   { return e.Error() }
func (e *errAlreadyExists) GRPCCode() codes.Code  { return codes.AlreadyExists }
func (e *errAlreadyExists) GRPCMessage() string   { return e.Error() }
