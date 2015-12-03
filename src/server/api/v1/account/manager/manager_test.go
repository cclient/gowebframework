package manager

import (
	"fmt"
	//	"gopkg.in/mgo.v2/bson"
	//	"server/common/tool"
	"testing"

	//	shopmanager "server/api/v1/shop/manager"
)

//
////
//func Test_GetAccountById(t *testing.T) {
//	res, err := GetAccountById("560b901a793513340237cad9")
//	fmt.Println(res)
//	fmt.Println(err)
//}
//
//func Test_GetAccounts(t *testing.T) {
//	res, err := GetAccounts(nil, 0, 0)
//	fmt.Println(res[0])
//	fmt.Println(len(res))
//	fmt.Println(err)
//}
//
//func Test_GetAccountsPage(t *testing.T) {
//	res, err := GetAccountsPage(tool.Meta{Limit: 100}, nil)
//	fmt.Println(res)
//	fmt.Println(err)
//}
//func Test_UpdateAccountById(t *testing.T) {
//	//修改
//	//	res, err := UpdateAccountById("560b901a793513340237cad9",bson.D{{"$set",bson.D{bson.DocElem{"userSource", 1}}}})
//	//	res, err := UpdateAccountById("560b901a793513340237cad9", bson.D{{"$set", bson.M{"help2": false}}})
//
//	//覆盖
//	res, err := UpdateAccountById("560b901a793513340237cad9", bson.D{bson.DocElem{"userSource3", 1}})
//	//	res, err := UpdateAccountById("560b901a793513340237cad9", &api.Meta{Offset: 3})
//	fmt.Println(res)
//	fmt.Println(err)
//}

//func Test_InsertAccount(t *testing.T) {
//	InsertAccount(Account{Isadmin: true,Truename:"hello word"})
//}

//func Test_InsertAccount(t *testing.T) {
//	//	account, _ := GetAccountById("55d54ac5c4456f8430c980a4")
//	//	account, _ := GetAccountById("55f11f9aa3536ec031541583")
//	account, _ := GetAccountById("55fbe2aa0f73146d41b59a0c")
//
//	getAccountManageShop(account)
//}
//
//func Test_login(t *testing.T) {
//	res, err := login("email@goyoo.com", "123456")
//	fmt.Println(res)
//	fmt.Println(err)
//
//}

func Test_GetAccountManageShopIds(t *testing.T) {
	res, err := Login("lixuelai@goyoo.com", "123456")
	fmt.Println(res)
	shopids, err := GetAccountManageShopIds(res)
	//	shopmanager.SetShopsAps(shops)
	fmt.Println(len(shopids))
	//	fmt.Println(shopids[0].Aps)
	fmt.Println(err)
	str := "14092821200001"
	fmt.Println(str)
	fmt.Println(str[0:10])
}
