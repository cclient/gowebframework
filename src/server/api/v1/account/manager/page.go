package manager

import (
	accountdao "server/api/v1/account/dao"
	accountmodel "server/api/v1/account/model"
	"server/common/tool"
)

type accountMeta struct {
	Meta tool.Meta              `json:"meta" bson:"meta"`
	List []accountmodel.Account `json:"list" bson:"list"`
}

func GetAccountsPage(meta tool.Meta, query interface{}) (*tool.ResponsePage, error) {
	session, _, collection, err := tool.GetCollection(nil, accountdao.DBNAME, accountdao.COLLECTIONNAME)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	if meta.Total == 0 {
		count, _ := collection.Find(query).Count()
		meta.Total = count
	}
	result, err := accountdao.GetAccounts(query, meta.Offset, meta.Limit)
	meta.Length = len(result)
	meta.SetRemaining()
	//	return &accountMeta{Meta: meta, List: result}, err
	return &tool.ResponsePage{Meta: meta, List: result}, err
}

//
//func GetAccountsPage(meta tool.Meta, query interface{}) (*tool.ResponsePage, error) {
//
//	arrcontain := []accountmodel.Account{}
//	vals := make([]interface{}, len(arrcontain))
//	for i, v := range arrcontain {
//		vals[i] = v
//	}
//	return tool.GetPage(DBNAME, COLLECTIONNAME, meta, query, vals)
//	//	session, _, collection, err := tool.GetCollection(nil, DBNAME, COLLECTIONNAME)
//	//	if err != nil {
//	//		return nil, err
//	//	}
//	//	defer session.Close()
//	//	if meta.Total == 0 {
//	//		count, _ := collection.Find(query).Count()
//	//		meta.Total = count
//	//	}
//	//	result, err := GetAccounts(query, meta.Offset, meta.Limit)
//	//	meta.Length = len(result)
//	//	meta.SetRemaining()
//	//	//	return &accountMeta{Meta: meta, List: result}, err
//	//	return &tool.ResponsePage{Meta: meta, List: result}, err
//}
