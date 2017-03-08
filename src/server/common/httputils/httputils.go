package httputils

import (
	"encoding/json"
	"errcode"
	"fmt"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"server/common/tool"
	"server/pkg/version"
	"strings"
	//	"utils"
)

// APIVersionKey is the client's requested API

const APIVersionKey = "api-version"

// APIFunc is an adapter to allow the use of ordinary functions as Docker API endpoints.
// Any function that has the appropriate signature can be register as a API endpoint (e.g. getVersion).
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error

// HijackConnection interrupts the http response writer to get the
// underlying connection and operate with it.
func HijackConnection(w http.ResponseWriter) (io.ReadCloser, io.Writer, error) {
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		return nil, nil, err
	}
	// Flush the options to make sure the client sets the raw mode
	conn.Write([]byte{})
	return conn, conn, nil
}

// CloseStreams ensures that a list for http streams are properly closed.
func CloseStreams(streams ...interface{}) {
	for _, stream := range streams {
		if tcpc, ok := stream.(interface {
			CloseWrite() error
		}); ok {
			tcpc.CloseWrite()
		} else if closer, ok := stream.(io.Closer); ok {
			closer.Close()
		}
	}
}

// CheckForJSON makes sure that the request's Content-Type is application/json.
func CheckForJSON(r *http.Request) error {
	ct := r.Header.Get("Content-Type")
	// No Content-Type header is ok as long as there's no Body
	if ct == "" {
		if r.Body == nil || r.ContentLength == 0 {
			return nil
		}
	}

	// Otherwise it better be json
	if tool.MatchesContentType(ct, "application/json") {
		return nil
	}
	return fmt.Errorf("Content-Type specified (%s) must be 'application/json'", ct)
}

// ParseForm ensures the request form is parsed even with invalid content types.
// If we don't do this, POST method without Content-type (even with empty body) will fail.
func ParseForm(r *http.Request) error {
	if r == nil {
		return nil
	}
	if err := r.ParseForm(); err != nil && !strings.HasPrefix(err.Error(), "mime:") {
		return err
	}
	return nil
}

// ParseMultipartForm ensure the request form is parsed, even with invalid content types.
func ParseMultipartForm(r *http.Request) error {
	if err := r.ParseMultipartForm(4096); err != nil && !strings.HasPrefix(err.Error(), "mime:") {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, err error) error {
	if err == nil || w == nil {
		logrus.WithFields(logrus.Fields{"error": err, "writer": w}).Error("unexpected HTTP error handling")
		return nil
	}
	var errinfo errcode.ErrorDescriptor
	switch err.(type) {
	case errcode.ErrorCode:
		daError, _ := err.(errcode.ErrorCode)
		errinfo = daError.Descriptor()
		return WriteJSON(w, errinfo.HTTPStatusCode, errinfo)
	case errcode.Error:
		daError, _ := err.(errcode.Error)
		errinfo = daError.ErrorCode().Descriptor()
		errinfo.Message = daError.Message
		return WriteJSON(w, errinfo.HTTPStatusCode, errinfo)
	default:
		errStr := strings.ToLower(err.Error())
		statusCode := http.StatusInternalServerError
		for keyword, status := range map[string]int{
			"not found":             http.StatusNotFound,
			"no such":               http.StatusNotFound,
			"bad parameter":         http.StatusBadRequest,
			"conflict":              http.StatusConflict,
			"impossible":            http.StatusNotAcceptable,
			"wrong login/password":  http.StatusUnauthorized,
			"hasn't been activated": http.StatusForbidden,
		} {
			if strings.Contains(errStr, keyword) {
				statusCode = status
				break
			}
		}
		http.Error(w, errStr, statusCode)
		return nil
	}
}

// WriteJSON writes the value v to the http response stream as json with standard json encoding.
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	fmt.Println(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func WriteSuccess(w http.ResponseWriter, code int) error {
	fmt.Println(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	item := Item{Item: "success"}
	return json.NewEncoder(w).Encode(item)
}

type Item struct {
	Item string `json:"item"`
}

func Redirect(w http.ResponseWriter, r *http.Request, urlStr string) error {
	http.Redirect(w, r, urlStr, 302)
	return nil
}

// VersionFromContext returns an API version from the context using APIVersionKey.
// It panics if the context value does not have version.Version type.
func VersionFromContext(ctx context.Context) (ver version.Version) {
	if ctx == nil {
		return
	}
	val := ctx.Value(APIVersionKey)
	if val == nil {
		return
	}
	return val.(version.Version)
}
