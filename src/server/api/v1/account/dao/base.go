package dao

import (
	"server/api/v1/account/model"
	"server/common/tool"
	"server/errors"
)

var COLLECTIONNAME = "account"
var DBNAME = "shenji"

func InsertAccount(account model.Account) error {
	//TODO 以后写个连接池
	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	defer session.Close()
	err = collection.Insert(&account)
	return err
}

func GetAccountById(id string) (account *model.Account, err error) {
	return GetAccountByMgoId(id, true)
}

func GetAccount(query interface{}) (account *model.Account, err error) {
	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return nil, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	defer session.Close()
	account = new(model.Account)
	err = collection.Find(query).One(&account)
	return account, err
}

func GetAccountByMgoId(id interface{}, ismgoid bool) (*model.Account, error) {
	id, err := tool.GetTrueQueryId(id, ismgoid)
	if err != nil {
		return nil, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return nil, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	account := new(model.Account)
	defer session.Close()
	err = collection.FindId(id).One(account)
	if err != nil {
		return nil, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	return account, nil
}

func UpdateAccountById(id string, update interface{}) (bool, error) {
	return UpdateAccountByMgoId(id, true, update)
}

func UpdateAccounts(selector interface{}, update interface{}) (int, error) {
	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return 0, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	defer session.Close()
	info, err := collection.UpdateAll(selector, update)
	if err != nil {
		return 0, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	return info.Updated, nil
}

func UpdateAccountByMgoId(id interface{}, ismgoid bool, update interface{}) (bool, error) {
	id, err := tool.GetTrueQueryId(id, ismgoid)
	if err != nil {
		return false, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return false, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	defer session.Close()
	err = collection.UpdateId(id, update)
	if err != nil {
		return false, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	return true, nil
}
func RemoveAccountById(id string) (bool, error) {
	return RemoveAccountByMgoId(id, true)
}
func RemoveAccounts(selector interface{}) (int, error) {
	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return 0, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	defer session.Close()
	info, err := collection.RemoveAll(selector)
	if err != nil {
		return 0, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	return info.Removed, err
}

func RemoveAccountByMgoId(id interface{}, ismgoid bool) (bool, error) {
	id, err := tool.GetTrueQueryId(id, ismgoid)
	if err != nil {
		return false, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return false, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	defer session.Close()
	err = collection.RemoveId(id)
	if err != nil {
		return false, errors.ErrorCodeMongoError.WithArgs(err.Error())
	}
	return true, nil
}

func GetAccounts(query interface{}, skip int, limit int) (accounts []model.Account, err error) {
	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	mquery := tool.MgoQuerySkipLimit(collection.Find(query), skip, limit)
	err = mquery.Iter().All(&accounts)
	return
}
