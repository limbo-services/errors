package errors

import (
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
)

// ConflictError should be used when the provided infomation is in conflict with already exiting information
type ConflictError interface {
	error
	IsConflict() bool
}

// IsConflict returns true when:
// - err is a ConflictError
// - err equals sql.ErrNoRows
// - err is a mysql error (1022, 1062, 1088, 1092, 1291, 1557, 1569, 1586)
func IsConflict(err error) bool {
	if err != nil {
		if x, _ := err.(*mysql.MySQLError); x != nil {
			switch x.Number {
			case 1022,
				1062,
				1088,
				1092,
				1291,
				1557,
				1569,
				1586:
				return true
			}
		}

		x, _ := err.(ConflictError)
		if x != nil && x.IsConflict() {
			return true
		}
	}
	return false
}

// Conflict makes a new ConflictError
func Conflict(a ...interface{}) error {
	return &errConflict{s: fmt.Sprint(a...)}
}

// Conflictf is the formatted version of Conflict
func Conflictf(format string, a ...interface{}) error {
	return &errConflict{s: fmt.Sprintf(format, a...)}
}

type errConflict struct{ s string }

func (e *errConflict) Error() string        { return e.s }
func (e *errConflict) IsConflict() bool     { return true }
func (e *errConflict) HTTPStatusCode() int  { return http.StatusConflict }
func (e *errConflict) HTTPMessage() string  { return e.Error() }
func (e *errConflict) GRPCCode() codes.Code { return codes.FailedPrecondition }
func (e *errConflict) GRPCMessage() string  { return e.Error() }
