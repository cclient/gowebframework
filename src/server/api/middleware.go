package api

import (
	"net/http"
	//	"runtime"
	//	"strings"
	//	"api"
	"server/common/httputils"
	"server/common/tool"
	//	"errors"
	"github.com/gorilla/sessions"
	//session支持用
	//"github.com/kidstuff/mongostore"
	"golang.org/x/net/context"
	//	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
)

// middleware is an adapter to allow the use of ordinary functions as Docker API filters.
// Any function that has the appropriate signature can be register as a middleware.
type middleware func(handler httputils.APIFunc) httputils.APIFunc

//loggingMiddleware logs each request when logging is enabled.
func (s *Server) loggingMiddleware(handler httputils.APIFunc) httputils.APIFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		if s.cfg.Logging {
			logrus.Infof("%s %s", r.Method, r.RequestURI)
		}
		return handler(ctx, w, r, vars)
	}
}

//

func ParseFormMiddleware(handler httputils.APIFunc) httputils.APIFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		r.ParseForm()
		return handler(ctx, w, r, vars)
	}
}

func checkIsLogin(session *sessions.Session, err error) bool {
	if session != nil && err == nil {
		if _, ok := session.Values["account"]; ok {
			if session.Values["account"] != nil {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
func CheckLoginMiddleware(handler httputils.APIFunc) httputils.APIFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		if r.URL.Path != "/api/v1/account/login" {
			websession, err := tool.GetWebSession(r)
			fmt.Println("session", websession)
			fmt.Println("sessionvalue", websession.Values)
			islogin := checkIsLogin(websession, err)
			if islogin == true {
				return handler(ctx, w, r, vars)
			} else {
				//未登陆路转到跟目录
				return httputils.Redirect(w, r, "/")
			}
		} else {
			return handler(ctx, w, r, vars)
		}
	}
}

//根据请求更新session及cookie过期时间 时间默认10分钟
func SessionMiddleware(handler httputils.APIFunc) httputils.APIFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		session, _ := tool.GetWebSession(r)
		session.Values["hello"] = "word"
		session.Save(r, w)
		sessions.Save(r, w)
		return handler(ctx, w, r, vars)
	}
}

// handleWithGlobalMiddlwares wraps the handler function for a request with
// the server's global middlewares. The order of the middlewares is backwards,
// meaning that the first in the list will be evaludated last.

// Example: handleWithGlobalMiddlewares(s.getContainersName)

//	s.loggingMiddleware(
//		s.userAgentMiddleware(
//			s.corsMiddleware(
//				versionMiddleware(s.getContainersName)
//			)
//		)
//	)
// )

func (s *Server) handleWithGlobalMiddlewares(handler httputils.APIFunc) httputils.APIFunc {
	middlewares := []middleware{
		//form转换，之后querystring form 可以通过相同的方式取值
		ParseFormMiddleware,
		//记录日志
		s.loggingMiddleware,
		//更新session时间
		//SessionMiddleware,
		//检查是否豆登陆
		//CheckLoginMiddleware,
	}
	h := handler
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
