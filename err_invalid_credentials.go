package errors

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// InvalidCredentialsError should be used when invalid credentials were provided
type InvalidCredentialsError interface {
	error
	IsInvalidCredentials() bool
}

// IsInvalidCredentials returns true when:
// - err is a InvalidCredentialsError
// - grpc.Code is codes.Unauthenticated
func IsInvalidCredentials(err error) bool {
	if err != nil {
		if grpc.Code(err) == codes.Unauthenticated {
			return true
		}

		x, _ := err.(InvalidCredentialsError)
		if x != nil && x.IsInvalidCredentials() {
			return true
		}
	}
	return false
}

// InvalidCredentials makes a new InvalidCredentialsError
func InvalidCredentials(a ...interface{}) error {
	return &errInvalidCredentials{s: fmt.Sprint(a...)}
}

// InvalidCredentialsf is the formatted version of InvalidCredentials
func InvalidCredentialsf(format string, a ...interface{}) error {
	return &errInvalidCredentials{s: fmt.Sprintf(format, a...)}
}

type errInvalidCredentials struct{ s string }

func (e *errInvalidCredentials) Error() string              { return e.s }
func (e *errInvalidCredentials) IsInvalidCredentials() bool { return true }
func (e *errInvalidCredentials) HTTPStatusCode() int        { return http.StatusUnauthorized }
func (e *errInvalidCredentials) HTTPMessage() string        { return e.Error() }
func (e *errInvalidCredentials) GRPCCode() codes.Code       { return codes.Unauthenticated }
func (e *errInvalidCredentials) GRPCMessage() string        { return e.Error() }
