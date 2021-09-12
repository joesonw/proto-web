package errutil

import (
	"net/http"

	"google.golang.org/grpc/status"
)

func IsTemporary(err error) bool {
	_, ok := err.(interface{
		Temporary()
	})
	return ok
}

func IsFatal(err error) bool {
	_, ok := err.(interface{
		Fatal()
	})
	return ok
}

func HTTPStatus(err error, or... int) int {
	hs, ok := err.(interface{
		HTTPStatus() int
	})
	if !ok {
		if len(or) > 0 {
			return or[0]
		}
		return http.StatusInternalServerError
	}
	return hs.HTTPStatus()
}

func GRPCStatus(err error) *status.Status {
	gs, ok := err.(interface{
		GRPCStatus() *status.Status
	})
	if !ok {
		return nil
	}
	return gs.GRPCStatus()
}