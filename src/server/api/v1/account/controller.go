package account

import (
	"crypto/md5"
	"fmt"
	"github.com/gorilla/schema"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	accountdao "server/api/v1/account/dao"
	accountmanager "server/api/v1/account/manager"
	accountmodel "server/api/v1/account/model"
	operalogdao "server/api/v1/operalog/dao"
	operalogmodel "server/api/v1/operalog/model"
	"server/common/httputils"
	"server/common/tool"
	"server/errors"
	"strconv"
	"strings"
)

func (s *accountRouter) GetAccountById(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	accountid := vars["id"]
	if accountid == "" {
		return httputils.WriteError(w, accountmanager.ErrorCodeInvalidAccountId)
	}
	res, err := accountdao.GetAccountById(accountid)
	if res != nil {
		return httputils.WriteJSON(w, http.StatusOK, tool.ResponseItem{Item: res})
	} else if err == nil {
		return httputils.WriteError(w, accountmanager.ErrorCodeIdGetNULL.WithArgs(accountid))
	} else {
		return httputils.WriteError(w, err)
	}
}

func (s *accountRouter) getAccountManagePCA(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	httputils.ParseForm(r)
	//查询场所条件
	qm := make(tool.RequestMongoKeyMap)
	issa, shopids, err := accountmanager.GetAccountManageShopIdsByAccountId(r.FormValue("accountid"))
	if err != nil {
		return httputils.WriteJSON(w, http.StatusOK, errors.ErrorCodeOther.WithArgs("get manager shop error"))
	} else {
		if issa != true {
			if shopids != nil && len(shopids) > 0 {
				qm.SetInStringArray("_id", shopids)
			} else {
				return httputils.WriteJSON(w, http.StatusOK, errors.ErrorCodeOther.WithArgs("get manager shop error"))
			}
		}
	}
	res, err := accountmanager.GetAccountManageShopSimplesPCA(qm.ParseRequestToMongoQuery())
	pcasmap := map[string]int{}
	rolecount := 0
	for _, shop := range res {
		arr := make([]string, 3)
		arr[0] = shop.Province
		arr[1] = shop.City
		arr[2] = shop.Area
		if shop.Province != "" {
			pcastr := strings.Join(arr, ",")
			if pcasmap[pcastr] != 1 {
				pcasmap[strings.Join(arr, ",")] = 1
				rolecount++
			}
		}
	}
	roles := make([]accountmodel.Role, rolecount)
	pcaarr := make([]string, 3)
	roleiter := 0
	for key, _ := range pcasmap {
		pcaarr = strings.Split(key, ",")
		roles[roleiter] = accountmodel.Role{Province: pcaarr[0], City: pcaarr[1], Area: pcaarr[2]}
		roleiter++
	}
	if res != nil {
		return httputils.WriteJSON(w, http.StatusOK, tool.ResponsePage{List: roles})
	} else if err != nil {
		return httputils.WriteError(w, errors.ErrorCodeOther.WithArgs(err.Error()))
	} else {
		return httputils.WriteJSON(w, http.StatusOK, tool.ResponsePage{List: nil})
	}
}

func CheckHasPermi(pcastring string, roles []accountmodel.Role) bool {
	for _, role := range roles {
		mk := []string{}
		if role.Province != "" {
			mk = append(mk, role.Province)
		}
		if role.City != "" {
			mk = append(mk, role.City)
		}
		if role.Area != "" {
			mk = append(mk, role.Area)
		}
		rolestring := strings.Join(mk, ",")
		if strings.Index(pcastring, rolestring) != -1 {
			return true
		} else {
			continue
		}
	}
	return false
}
func (s *accountRouter) GetAccountsPage(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	httputils.ParseForm(r)
	meta := httputils.GetPagination(r)
	qm := make(tool.RequestMongoKeyMap)
	qm.SetMgoQKey("truename", r.FormValue("truename"))
	qm.SetMgoQKey("email", r.FormValue("email"))
	qm.SetMgoQKey("name", r.FormValue("name"))
	qm.SetMgoQKey("telphone", r.FormValue("telphone"))
	accountid := r.FormValue("accountid")
	account, err := accountdao.GetAccountById(accountid)
	fmt.Println(account, err)
	manageraccountids := []bson.ObjectId{}
	if err != nil {
		return httputils.WriteError(w, errors.ErrorCodeOther.WithArgs(err.Error()))
	}
	if account.IsSuperAdmin == false {
		ajrs, _ := accountmanager.GetAccountJustRoles(nil, 0, 0)
		for _, ajr := range ajrs {
			for _, role := range ajr.Roles {
				str := role.ToPCAString()
				if CheckHasPermi(str, account.Roles) == true {
					fmt.Println(str)
					manageraccountids = append(manageraccountids, ajr.Id)
					break
				}
			}
		}
		qm.SetInMgoIdArray("_id", manageraccountids)
	}
	res, err := accountmanager.GetAccountsPage(meta, qm.ParseRequestToMongoQuery())
	if res != nil {
		return httputils.WriteJSON(w, http.StatusOK, res)
	} else if err == nil {
		return httputils.WriteError(w, errors.ErrorCodeOther.WithArgs(err.Error()))
	} else {
		return httputils.WriteError(w, err)
	}
}

func getRoles(r *http.Request) []accountmodel.Role {
	rolekeys := []string{"province", "city", "area"}
	rolemaxcount := 5
	rolenumcount := 0
	fmt.Println(r.Form)
	fmt.Println("roles[" + strconv.Itoa(1) + "]" + "[" + rolekeys[0] + "]")
	fmt.Println(r.FormValue("roles[" + strconv.Itoa(1) + "]" + "[" + rolekeys[0] + "]"))
	for i := 0; i < rolemaxcount; i++ {
		keyper := "roles[" + strconv.Itoa(i) + "]"
		if r.FormValue(keyper+"["+rolekeys[0]+"]") != "" {
			rolenumcount++
		} else {
			break
		}
	}
	roles := make([]accountmodel.Role, rolenumcount)
	for i := 0; i < rolenumcount; i++ {
		keyper := "roles[" + strconv.Itoa(i) + "]"
		fmt.Println(keyper+"["+rolekeys[0]+"]", r.FormValue(keyper+"["+rolekeys[0]+"]"))
		role := accountmodel.Role{
			Province: r.FormValue(keyper + "[" + rolekeys[0] + "]"),
			City:     r.FormValue(keyper + "[" + rolekeys[1] + "]"),
			Area:     r.FormValue(keyper + "[" + rolekeys[2] + "]"),
		}
		roles[i] = role
	}
	return roles
}
func (s *accountRouter) InsertAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	httputils.ParseForm(r)
	account := accountmodel.Account{}
	decoder := schema.NewDecoder()
	err := decoder.Decode(&account, r.Form)
	account.Roles = getRoles(r)
	account.Id = tool.CreateMgoId()
	fmt.Println(account.PasswdNew)
	account.PasswdNew = fmt.Sprintf("%x", md5.Sum([]byte(account.PasswdNew)))
	err = accountmanager.InsertAccount(account)
	if err == nil {
		operaaccount, _ := accountdao.GetAccountById(r.FormValue("accountid"))
		operalogdao.InsertAccountOperalog(operaaccount, operalogmodel.AccountAdd, nil, account)
		return httputils.WriteSuccess(w, http.StatusOK)
	} else {
		return httputils.WriteError(w, accountmanager.ErrorCodeUpdateFAILD)
	}

}

func (s *accountRouter) UpdateAccountById(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	res, _ := accountmanager.UpdateAccountById("", nil)
	return httputils.WriteJSON(w, http.StatusOK, res)
}

func (s *accountRouter) RemoveAccountById(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	httputils.ParseForm(r)
	id := vars["id"]
	if id == "" {
		id = r.FormValue("id")
	}
	res, err := accountmanager.RemoveAccountById(id)
	fmt.Println(err)
	if err == nil && res == true {
		operaaccount, _ := accountdao.GetAccountById(r.FormValue("accountid"))
		operalogdao.InsertAccountOperalog(operaaccount, operalogmodel.AccountDel, id, nil)
		return httputils.WriteSuccess(w, http.StatusOK)
	}
	return httputils.WriteError(w, accountmanager.ErrorCodeRemoveFAILD)
}
func (s *accountRouter) UpdateAccountInfo(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	httputils.ParseForm(r)
	//不能直接接收一个json 对象，比较悲剧
	id := r.FormValue("id")
	qm := make(tool.RequestMongoKeyMap)
	kmaps := tool.KeyMaps{
		{MgoKey: "email"},
		{MgoKey: "telphone"},
		{MgoKey: "truename"},
		{MgoKey: "noticemessage"},
		{MgoKey: "noticeemail"},
	}
	kmaps.SetReqKeys()
	for _, kmap := range kmaps {
		qm.SetMgoQKey(kmap.MgoKey, r.FormValue(kmap.ReqKey))
	}
	qm.SetMgoQKeyBool("isnoticeemail", r.FormValue("isnoticeemail"), httputils.BoolValueOrDefault(r, "isnoticeemail", false))
	qm.SetMgoQKeyBool("isnoticemessage", r.FormValue("isnoticemessage"), httputils.BoolValueOrDefault(r, "isnoticemessage", false))
	qm.SetMgoQKeyValue("roles", getRoles(r))
	fmt.Println(qm.ParseUpdateToMongoSet())
	res, err := accountmanager.UpdateAccountById(id, qm.ParseUpdateToMongoSet())
	fmt.Println(res, err)
	if err == nil && res == true {
		operaaccount, _ := accountdao.GetAccountById(r.FormValue("accountid"))
		operalogdao.InsertAccountOperalog(operaaccount, operalogmodel.AccountUpdate, id, qm.ParseUpdateToMongoSet())
		return httputils.WriteSuccess(w, http.StatusOK)
	} else {
		return httputils.WriteError(w, accountmanager.ErrorCodeUpdateFAILD)
	}
}
