package tool

import (
	"fmt"
	"mime"
	"net/http"
	//	"path/filepath"
	//	"sort"
	//	"strconv"
	"github.com/Sirupsen/logrus"
	"strings"
	//	"server/pkg/system"
	"github.com/gorilla/sessions"
	//	"github.com/kidstuff/mongostore"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

// byPrivatePort is temporary type used to sort types.Port by PrivatePort
//type byPrivatePort []types.Port

//func (r byPrivatePort) Len() int           { return len(r) }
//func (r byPrivatePort) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
//func (r byPrivatePort) Less(i, j int) bool { return r[i].PrivatePort < r[j].PrivatePort }

// MatchesContentType validates the content type against the expected one
func MatchesContentType(contentType, expectedType string) bool {
	mimetype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		logrus.Errorf("Error parsing media type: %s error: %v", contentType, err)
	}
	return err == nil && mimetype == expectedType
}

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)

	fmt.Println(path)
	splitstring := strings.Split(path, "\\")
	size := len(splitstring)
	splitstring = strings.Split(path, splitstring[size-1])
	ret := strings.Replace(splitstring[0], "\\", "/", size-1)
	return ret
}

func GetPage(DBNAME string, COLLECTIONNAME string, meta Meta, query interface{}, gettypeinfo interface{}) (*ResponsePage, error) {
	var arrcontain []interface{}
	session, _, collection, err := GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	if meta.Total == 0 {
		count, _ := collection.Find(query).Count()
		meta.Total = count
	}
	mquery := MgoQuerySkipLimit(collection.Find(query), meta.Offset, meta.Limit)
	err = mquery.Iter().All(&arrcontain)
	meta.Length = len(arrcontain)
	meta.SetRemaining()
	return &ResponsePage{Meta: meta, List: arrcontain}, err
}
func MgoIdToString(id bson.ObjectId) string {
	return fmt.Sprintf("%x", string(id))
}

func MgoQuerySkipLimit(query *mgo.Query, skip int, limit int) *mgo.Query {
	if skip != 0 {
		query = query.Skip(skip)
	}
	if limit != 0 {
		query = query.Limit(limit)
	}
	return query
}

func GetWebSession(r *http.Request) (*sessions.Session, error) {
	var store = sessions.NewCookieStore([]byte("something-very-secret"))
	websession, err := store.Get(r, "session-key")
	fmt.Println("fun ", websession.Values)
	return websession, err
}

func GetTodayDay2330() time.Time {
	return GetBeforeDay2330(0)
}

func GetNowTs() int64 {
	return time.Now().Unix()
}

func GetTimeShortString(t time.Time) string {
	m := strconv.Itoa(int(t.Month()))
	if len(m) == 1 {
		m = "0" + m
	}
	d := strconv.Itoa(t.Day())
	if len(d) == 1 {
		d = "0" + m
	}
	return (strconv.Itoa(t.Year()))[2:4] + m + d
}

func GetYesteryDay2330() time.Time {
	return GetBeforeDay2330(1)
}
func GetBeforeDay2330(beforenum int) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day()-beforenum, 23, 30, 0, 0, now.Location())

}

func MacToApMac(mac string) string {
	mac = strings.Replace(mac, ":", "", -1)
	mac = strings.Replace(mac, "-", "", -1)
	mk := make([]string, 6)
	j := 0
	for i := 0; i <= len(mac); i++ {
		if i != 0 && i%2 == 0 {
			//			fmt.Println(mac[i-2 : i])
			mk[j] = mac[i-2 : i]
			j++
		}
	}
	return strings.ToUpper(strings.Join(mk, "-"))
}

//11分开始执行  设为10的时间
//取10分相隔的时间
func GetTenMiniute() time.Time {
	now := time.Now()
	minute := now.Minute() - now.Minute()%10
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), minute, 0, 0, now.Location())
}
