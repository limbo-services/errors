package errors

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

// MissingCredentialsError should be used when no credentials were found
type MissingCredentialsError interface {
	error
	IsMissingCredentials() bool
}

// IsMissingCredentials returns true when err is a MissingCredentialsError
func IsMissingCredentials(err error) bool {
	if err != nil {
		x, _ := err.(MissingCredentialsError)
		if x != nil && x.IsMissingCredentials() {
			return true
		}
	}
	return false
}

// MissingCredentials makes a new MissingCredentialsError
func MissingCredentials(a ...interface{}) error {
	return &errMissingCredentials{s: fmt.Sprint(a...)}
}

// MissingCredentialsf is the formatted version of MissingCredentials
func MissingCredentialsf(format string, a ...interface{}) error {
	return &errMissingCredentials{s: fmt.Sprintf(format, a...)}
}

type errMissingCredentials struct{ s string }

func (e *errMissingCredentials) Error() string              { return e.s }
func (e *errMissingCredentials) IsMissingCredentials() bool { return true }
func (e *errMissingCredentials) HTTPStatusCode() int        { return http.StatusUnauthorized }
func (e *errMissingCredentials) HTTPMessage() string        { return e.Error() }
func (e *errMissingCredentials) GRPCCode() codes.Code       { return codes.Unauthenticated }
func (e *errMissingCredentials) GRPCMessage() string        { return e.Error() }
