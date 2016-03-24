package errors

import (
	"fmt"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// PermissionDeniedError should be used in case of a permission error
type PermissionDeniedError interface {
	error
	IsPermissionDenied() bool
}

// IsPermissionDenied returns true when:
// - err is a PermissionDeniedError
// - os.IsPermission is true
// - grpc.Code is codes.PermissionDenied
func IsPermissionDenied(err error) bool {
	if err != nil {
		if os.IsPermission(err) {
			return true
		}
		if grpc.Code(err) == codes.PermissionDenied {
			return true
		}
		x, _ := err.(PermissionDeniedError)
		if x != nil && x.IsPermissionDenied() {
			return true
		}
	}
	return false
}

// PermissionDenied makes a new PermissionDeniedError
func PermissionDenied(a ...interface{}) error {
	return &errPermissionDenied{s: fmt.Sprint(a...)}
}

// PermissionDeniedf is the formatted version of PermissionDenied
func PermissionDeniedf(format string, a ...interface{}) error {
	return &errPermissionDenied{s: fmt.Sprintf(format, a...)}
}

type errPermissionDenied struct{ s string }

func (e *errPermissionDenied) Error() string            { return e.s }
func (e *errPermissionDenied) IsPermissionDenied() bool { return true }
func (e *errPermissionDenied) HTTPStatusCode() int      { return http.StatusForbidden }
func (e *errPermissionDenied) HTTPMessage() string      { return e.Error() }
func (e *errPermissionDenied) GRPCCode() codes.Code     { return codes.PermissionDenied }
func (e *errPermissionDenied) GRPCMessage() string      { return e.Error() }
