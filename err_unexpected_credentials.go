package errors

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

// UnexpectedCredentialsError should be used when no (or different) credentials were expected
type UnexpectedCredentialsError interface {
	error
	IsUnexpectedCredentials() bool
}

// IsUnexpectedCredentials returns true when err is a UnexpectedCredentialsError
func IsUnexpectedCredentials(err error) bool {
	if err != nil {
		x, _ := err.(UnexpectedCredentialsError)
		if x != nil && x.IsUnexpectedCredentials() {
			return true
		}
	}
	return false
}

// UnexpectedCredentials makes a new UnexpectedCredentialsError
func UnexpectedCredentials(a ...interface{}) error {
	return &errUnexpectedCredentials{s: fmt.Sprint(a...)}
}

// UnexpectedCredentialsf is the formatted version of UnexpectedCredentials
func UnexpectedCredentialsf(format string, a ...interface{}) error {
	return &errUnexpectedCredentials{s: fmt.Sprintf(format, a...)}
}

type errUnexpectedCredentials struct{ s string }

func (e *errUnexpectedCredentials) Error() string                 { return e.s }
func (e *errUnexpectedCredentials) IsUnexpectedCredentials() bool { return true }
func (e *errUnexpectedCredentials) HTTPStatusCode() int           { return http.StatusUnauthorized }
func (e *errUnexpectedCredentials) HTTPMessage() string           { return e.Error() }
func (e *errUnexpectedCredentials) GRPCCode() codes.Code          { return codes.Unauthenticated }
func (e *errUnexpectedCredentials) GRPCMessage() string           { return e.Error() }
