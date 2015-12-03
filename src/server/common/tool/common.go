package tool

import (
	"fmt"
	"mime"
	"net/http"
	//	"path/filepath"
	//	"sort"
	//	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	//	"server/pkg/system"
	"github.com/gorilla/sessions"
	//	"github.com/kidstuff/mongostore"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"os/exec"
	"path/filepath"
	//	"server/errors"
	"server/pkg/version"
	"strconv"
	"time"
	//	accountmodel "server/api/v1/account/model"
)

// Common constants for daemon and client.
const (
	// Version of Current REST API
	Version version.Version = "1.22"

	// MinVersion represents Minimun REST API version supported
	MinVersion version.Version = "1.12"

	// DefaultDockerfileName is the Default filename with Docker commands, read by docker build
	DefaultDockerfileName string = "Dockerfile"
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

//func getOpearStringArray(opera string, strarr []string) bson.M {
//	if len(strarr) != 0 {
//		return bson.M{opera: strarr}
//	}
//	return nil
//}
//

//func GetPageQueryAndSetMeta(DBNAME string, COLLECTIONNAME string, meta *Meta, query interface{}, gettypeinfo interface{})  {
//	session, _, collection, err := GetCollection(nil, DBNAME, COLLECTIONNAME)
//	if err != nil {
//		return nil, err
//	}
//	defer session.Close()
//	if meta.Total == 0 {
//		count, _ := collection.Find(query).Count()
//		meta.Total = count
//	}
//	return
//}

//func GetPageData(collection mgo.Collection,query interface{}, skip int, limit int) (aps []apmodel.Ap, err error) {
//	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
//	if err != nil {
//		return nil, err
//	}
//	defer session.Close()
//	mquery := tool.MgoQuerySkipLimit(collection.Find(query), skip, limit)
//	err = mquery.Iter().All(&aps)
//	return
//}

//func GetPage(DBNAME string, COLLECTIONNAME string, meta Meta, query interface{}, gettypeinfo interface{}) (*ResponsePage, error) {
//
//	var arrcontain []interface{}
//	switch gettypeinfo.(type) { //多选语句switch
//
//	case accountmodel.Account:
//		arrcontain = []accountmodel.Account{}
//		//是字符时做的事情
//	}
//
//	//	session, _, collection, err := tool.GetCollection(nil, "shenji", "ap")
//	session, _, collection, err := GetCollection(nil, DBNAME, COLLECTIONNAME)
//	if err != nil {
//		return nil, err
//	}
//	defer session.Close()
//	if meta.Total == 0 {
//		count, _ := collection.Find(query).Count()
//		meta.Total = count
//	}
//	mquery := MgoQuerySkipLimit(collection.Find(query), meta.Offset, meta.Limit)
//	err = mquery.Iter().All(&arrcontain)
//	meta.Length = len(arrcontain)
//	meta.SetRemaining()
//	return &ResponsePage{Meta: meta, List: arrcontain}, err
//}
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
	//	_, _, collection, err := GetCollection(nil, "sessiondb", "session")
	//	if err == nil {
	var store = sessions.NewCookieStore([]byte("something-very-secret"))
	//	store := mongostore.NewMongoStore(collection, 6000, true, []byte("secret-key-goyoo"))
	websession, err := store.Get(r, "session-key")
	fmt.Println("fun ", websession.Values)
	return websession, err
	//	}
	//	return nil, err
}

func GetTodayDay2330() time.Time {
	return GetBeforeDay2330(0)
}
func GetNowTs() int64 {
	return time.Now().Unix()
}

//todo test
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
//function getTenMiniute() {
//    let now = new Date()
//    now.setMinutes(now.getMinutes() - now.getMinutes() % 10)
//    now.setSeconds(0)
//    now.setMilliseconds(0)
//    return now
//}
func GetTenMiniute() time.Time {
	now := time.Now()
	minute := now.Minute() - now.Minute()%10
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), minute, 0, 0, now.Location())
}

//type Pagination struct{
//	}
//pagination
//module.exports.getFromReq = function (query, defaultLimit)
//{
//    var offset = parseInt(query.offset, 10);
//    if (true === isNaN(offset) || 0 > offset) {
//        offset = 0;
//    }
//
//    var limit = parseInt(query.limit, 10);
//    if (true === isNaN(limit) ||
//        limit <= 0 ||
//        limit > defaultLimit) {
//        limit = defaultLimit;
//    }
//
//    return {
//        offset: offset,
//        limit: limit
//    };
//};

//
//module.exports.getFromReq = function (query, defaultLimit)
//{
//    var offset = parseInt(query.offset, 10);
//    if (true === isNaN(offset) || 0 > offset) {
//        offset = 0;
//    }
//
//    var limit = parseInt(query.limit, 10);
//    if (true === isNaN(limit) ||
//        limit <= 0 ||
//        limit > defaultLimit) {
//        limit = defaultLimit;
//    }
//
//    return {
//        offset: offset,
//        limit: limit
//    };
//};

//func (p Pagination) getMeta(totalLength int,currentLength int ){
//
////	module.exports.getMeta = function (totalLength, currentLength, pagination)
////{
////    return {
////        offset: pagination.offset,
////        limit: pagination.limit,
////        total: totalLength,
////        length: currentLength,
////        remaining: currentLength === 0 ? 0 : totalLength - pagination.offset - currentLength
////    };
////};
//
//	}
