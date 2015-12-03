package api

import (
	"net/http"
	//	"runtime"
	//	"strings"

	//	"api"
	"server/common/httputils"
	"server/common/tool"
	//	"github.com/Sirupsen/logrus"
	//	"errors"
	"github.com/gorilla/sessions"
	//	"github.com/kidstuff/mongostore"
	"golang.org/x/net/context"
	//	"encoding/json"
	//	"server/pkg/version"

	"fmt"
//	accountdao "server/api/v1/account/dao"
	//	"gopkg.in/mgo.v2"
)

// middleware is an adapter to allow the use of ordinary functions as Docker API filters.
// Any function that has the appropriate signature can be register as a middleware.
type middleware func(handler httputils.APIFunc) httputils.APIFunc

// loggingMiddleware logs each request when logging is enabled.
//func (s *Server) loggingMiddleware(handler httputils.APIFunc) httputils.APIFunc {
//	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
//		if s.cfg.Logging {
//			logrus.Infof("%s %s", r.Method, r.RequestURI)
//		}
//		return handler(ctx, w, r, vars)
//	}
//}
//

//parseform转换，之后querystring form 可以通过相同的方式取值
//func ParseFormMiddleware(handler httputils.APIFunc) httputils.APIFunc {
//	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
//
//		r.ParseForm()
//		return handler(ctx, w, r, vars)
//	}
//}

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
				return httputils.Redirect(w, r, "/")
				//				return handler(ctx, w, r, vars)
				//								   return func(s *accountRouter) RemoveAccountById(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
				//									res, _ := accountmanager.RemoveAccountById("")
				//									return httputils.WriteJSON(w, http.StatusOK, res)
				//								}
			}

			//			dbsession, _, collection, err := tool.GetCollection(nil, "test", "test_session")
			//			if err != nil {
			//				store := mongostore.NewMongoStore(collection, 600, true, []byte("secret-key"))
			//				store.Get(r, "session-key")
			//				websession, err := store.Get(r, "session-key")
			//
			//				//				sessions.Save(r, w)
			//			}
			//检查 session
		} else {
			return handler(ctx, w, r, vars)

		}
	}
}

//
//session, err := tool.GetWebSession(r)
//
////	var store = sessions.NewCookieStore([]byte("something-very-secret"))
////	//		store := mongostore.NewMongoStore(collection, 600, true, []byte("secret-key-goyoo"))
////	session, err = store.Get(r, "session-key")
////	fmt.Println("fun ", websession.Values)
//
//
//
//	if err != nil {
//		return httputils.WriteJSON(w, http.StatusOK, errors.ErrorCodeOther.WithArgs(err))
//	}
//	if session != nil {
//		session.Values["account"] = account
//
//		session.Values["foo"] = "bar"
//		fmt.Println("session ", session)
//		fmt.Println("save session value ", session.Values)
//		session.Save(r, w)
//	}
//
//
//根据请求更新session及cookie过期时间 时间默认10分钟
func SessionMiddleware(handler httputils.APIFunc) httputils.APIFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		//		session, _ := tool.GetWebSession(r)
		//		session.Values["hello"] = "word"
		//		session.Save(r,w)
		//		sessions.Save(r, w)
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

//func AccountMiddleware(handler httputils.APIFunc) httputils.APIFunc {
//	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
//		r.ParseForm()
//		accountid := r.FormValue("accountid")
//		account := accountdao.GetAccountById(accountid)
//		//		session, _ := tool.GetWebSession(r)
//		//		session.Values["hello"] = "word"
//		//		session.Save(r,w)
//		//		sessions.Save(r, w)
//		return handler(ctx, w, r, vars)
//	}
//}

func (s *Server) handleWithGlobalMiddlewares(handler httputils.APIFunc) httputils.APIFunc {
	middlewares := []middleware{
	//		CheckLoginMiddleware,
	//		ParseFormMiddleware,

	//		s.corsMiddleware,
	//		s.userAgentMiddleware,
	//		s.loggingMiddleware,
	//		SessionMiddleware,
	//		CheckLoginMiddleware,
	}

	h := handler
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
