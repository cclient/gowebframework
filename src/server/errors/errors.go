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
	
	

//	// ErrorCodeUnsupported is returned when an operation is not supported.
//	ErrorCodeUnsupported = Register("mongo", 2002, ErrorDescriptor{
//		Value:   "UNSUPPORTED",
//		Message: "The operation is unsupported.",
//		Description: `The operation was unsupported due to a missing
//		implementation or invalid set of parameters.`,
//		HTTPStatusCode: http.StatusMethodNotAllowed,
//	})
//
//	// ErrorCodeUnauthorized is returned if a request requires
//	// authentication.
//	ErrorCodeUnauthorized = Register("mongo", 2003, ErrorDescriptor{
//		Value:   "UNAUTHORIZED",
//		Message: "authentication required",
//		Description: `The access controller was unable to authenticate
//		the client. Often this will be accompanied by a
//		Www-Authenticate HTTP response header indicating how to
//		authenticate.`,
//		HTTPStatusCode: http.StatusUnauthorized,
//	})
//
//	// ErrorCodeDenied is returned if a client does not have sufficient
//	// permission to perform an action.
//	ErrorCodeDenied = Register("mongo", 1003, ErrorDescriptor{
//		Value:   "DENIED",
//		Message: "requested access to the resource is denied",
//		Description: `The access controller denied access for the
//		operation on a resource.`,
//		HTTPStatusCode: http.StatusForbidden,
//	})
//
//	// ErrorCodeUnavailable provides a common error to report unavialability
//	// of a service or endpoint.
//	ErrorCodeUnavailable = Register("mongo", 1004, ErrorDescriptor{
//		Value:          "UNAVAILABLE",
//		Message:        "service unavailable",
//		Description:    "Returned when a service is not available",
//		HTTPStatusCode: http.StatusServiceUnavailable,
//	})
)
