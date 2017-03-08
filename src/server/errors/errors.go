package errors

import (
	"errcode"
	"net/http"
)

var (
	// ErrorCodeUnknown is a generic error that can be used as a last
	// resort if there is no situation-specific error message that can be used
	ErrorCodeMongoError = errcode.Register("mongo", 30201, errcode.ErrorDescriptor{
		Value:          "MongoError",
		Message:        "'%s'",
		Description:    ``,
		HTTPStatusCode: http.StatusBadRequest,
	})

	ErrorCodeOther = errcode.Register("other", 40201, errcode.ErrorDescriptor{
		Value:          "OtherError",
		Message:        "'%s'",
		Description:    ``,
		HTTPStatusCode: http.StatusBadRequest,
	})
)
