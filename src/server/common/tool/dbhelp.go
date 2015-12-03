package tool

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"server/errors"
	"strconv"
	"fmt"
)

type KeyMaps []KeyMap

func (kmaps KeyMaps) SetReqKeys() {
	for i, _ := range kmaps {
		if kmaps[i].ReqKey == "" {
			kmaps[i].ReqKey = kmaps[i].MgoKey
		}
	}
}

type KeyMap struct {
	MgoKey string
	ReqKey string
}

type RequestMongoKeyMap map[string]interface{}

//todo test
func GetTimeSection(begintime int64, endtime int64) bson.M {
	//	fmt.Println(begintime)
	//	fmt.Println(endtime)
	if begintime != 0 || endtime != 0 {
		var bsonarr bson.D
		if begintime != 0 {
			//			fmt.Println("add starttime")
			bsonarr = append(bsonarr, bson.DocElem{"$gt", begintime})
		}
		if endtime != 0 {
			bsonarr = append(bsonarr, bson.DocElem{"$lt", endtime})
		}
		//		keymap[key]=bson.M{"$gt": begintime,""}
		return bsonarr.Map()
		//		(*keymap)[key] = bsonarr.Map()
	}
	//	fmt.Println("time nil")
	return nil
}

func getOpearStringArray(opera string, strarr []string) bson.M {
	if len(strarr) != 0 {
		return bson.M{opera: strarr}
	}
	return nil
}

func GetInStringArray(strarr []string) bson.M {
	return getOpearStringArray("$in", strarr)
}

func GetMgoQueryReg(reg string) *bson.RegEx {
	if reg != "" {
		return &bson.RegEx{Pattern: reg, Options: "i"}
	}
	return nil
}

func (keymap RequestMongoKeyMap) HasKey() bool {
	for _, _ = range keymap {
		return true
	}
	return false
}

//func getMo
func (keymap RequestMongoKeyMap) ParseRequestToMongoQuery() bson.M {
	var bsonarr bson.D
	for key, value := range keymap {
		//		fmt.Println("key")
		//		fmt.Println(value)
		if value != nil && value != "" && value != 0 {
			//			fmt.Println(key)
			//			fmt.Println(value)
			bsonarr = append(bsonarr, bson.DocElem{key, value})
		}
	}
	//	fmt.Println(bsonarr.Map())
	return bsonarr.Map()
}

func (keymap RequestMongoKeyMap) ParseUpdateToMongoSet() bson.D {
	var bsonarr bson.D
	for key, value := range keymap {
		if value != nil && value != "" && value != 0 {
			bsonarr = append(bsonarr, bson.DocElem{key, value})
		}
	}
	return bson.D{{"$set", bsonarr}}
}


func CreateMgoId() bson.ObjectId {
	return bson.NewObjectId()
}
func (keymap RequestMongoKeyMap) SetMgoQKey(key string, value string) {
	if value != "" {
		keymap[key] = value
	}
}
func (keymap RequestMongoKeyMap) SetMgoQKeyBool(key string, value string, truevalue bool) {
	if value != "" {
		fmt.Println("truevalue")
		fmt.Println(truevalue)
		keymap[key] = truevalue
	}
}

func (keymap RequestMongoKeyMap) SetMgoQKeyStringToInt(key string, value string) {
	if value != "" {
		val, err := strconv.Atoi(value)
		if err == nil {
			keymap[key] = val
		}

	}
}

func (keymap RequestMongoKeyMap) SetMgoQKeyInt(key string, value string, truevalue int) {
	if value != "" {
		fmt.Println("truevalue")
		fmt.Println(truevalue)
		keymap[key] = truevalue
	}
}

func (keymap RequestMongoKeyMap) SetMgoQKeyValue(key string, value interface{}) {
	if value != nil {
		keymap[key] = value
	}
}

func (keymap RequestMongoKeyMap) SetMgoQueryReg(key string, reg string) {
	if reg != "" {
		keymap[key] = bson.RegEx{Pattern: reg, Options: "i"}
	}
}

func (keymap RequestMongoKeyMap) SetOrArray(key string, strarr []string) {
	keymap.setOpearStringArray(key, "$or", strarr)
}
//func (keymap RequestMongoKeyMap) SetOrArray(key string, strarr []string) {
//	keymap.setOpearStringArray(key, "$or", strarr)
//}


func (keymap RequestMongoKeyMap) SetInStringArray(key string, strarr []string) {
	keymap.setOpearStringArray(key, "$in", strarr)
}
func (keymap RequestMongoKeyMap) SetInMgoIdArray(key string, strarr []bson.ObjectId) {
	keymap.setOpearMgoIdArray(key, "$in", strarr)
}
func (keymap RequestMongoKeyMap) setOpearMgoIdArray(key string, opera string, strarr []bson.ObjectId) {
	if strarr != nil && len(strarr) != 0 {
		keymap[key] = bson.M{opera: strarr}
	}
}

func (keymap RequestMongoKeyMap) SetInArray(key string, strarr []interface{}) {
	keymap.setOpearArray(key, "$in", strarr)
}

func (keymap RequestMongoKeyMap) setOpearArray(key string, opera string, strarr []interface{}) {
	if strarr != nil && len(strarr) != 0 {
		keymap[key] = bson.M{opera: strarr}
	}
}

func (keymap RequestMongoKeyMap) setOpearStringArray(key string, opera string, strarr []string) {
	if strarr != nil && len(strarr) != 0 {
		keymap[key] = bson.M{opera: strarr}
	}
}


func (keymap RequestMongoKeyMap) SetOpear(key string, opera string, value interface{}) {
	keymap[key] = bson.M{opera: value}
}

//func (keymap RequestMongoKeyMap) SetOpearDom(opera string, value interface{}) {
//	
//
//}

func (keymap RequestMongoKeyMap) SetTimeSection(key string, begintime int64, endtime int64) {
	if begintime != 0 || endtime != 0 {
		var bsonarr bson.D
		if begintime != 0 {
			bsonarr = append(bsonarr, bson.DocElem{"$gt", begintime})
		}
		if endtime != 0 {
			bsonarr = append(bsonarr, bson.DocElem{"$lt", endtime})
		}
		keymap[key] = bsonarr.Map()
		//		return bsonarr.Map()

		//		(*keymap)[key] = bsonarr.Map()
	}
	//	fmt.Println("time nil")
	//	return nil
}

var Port = "27017"
var AESKEY=[]byte("goyooxiaoyunaudi")

func GetCollection(oldsession *mgo.Session, db string, collectionname string) (session *mgo.Session, database *mgo.Database, collection *mgo.Collection, err error) {
	if oldsession == nil {
		session, err = mgo.Dial("127.0.0.1:" + Port)
		if err != nil {
			return nil, nil, nil, err
		}
		session.SetMode(mgo.Monotonic, true)
	} else {
		session = oldsession
	}
	database = session.DB(db)
	collection = database.C(collectionname)
	return
}

func GetDB(db string) (session *mgo.Session, database *mgo.Database, err error) {
	session, err = mgo.Dial("127.0.0.1:" + Port)
	if err != nil {
		return nil, nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	database = session.DB(db)
	return
}

func GetTrueQueryId(id interface{}, ismgoid bool) (interface{}, error) {
	if ismgoid == true {
		switch id.(type) {
		case string:
			mgoidstring := id.(string)
			if bson.IsObjectIdHex(mgoidstring) {
				id = bson.ObjectIdHex(mgoidstring)
			} else {
				return nil, errors.ErrorCodeMongoError.WithArgs("Invalid input to ObjectIdHex")
			}
		default:
		}
	}
	return id, nil
}

func GetMongoCollectionDataCount(dbname string, collectionname string, query interface{}) (int, error) {
	//TODO 以后写个连接池
	session, _, collection, err := GetCollection(nil, dbname, collectionname)
	var count int
	if err != nil {
		return 0, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	defer session.Close()
	if query != nil {
		count, err = collection.Find(query).Count()
	} else {
		count, err = collection.Count()
	}
	if err != nil {
		return 0, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	return count, nil
}
