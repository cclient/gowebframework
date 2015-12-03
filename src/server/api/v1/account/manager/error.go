package manager

import (
	"errcode"
	"net/http"
)


const errGroup = "account"
var (
	ErrorCodeInvalidAccountId = errcode.Register(errGroup, 20001, errcode.ErrorDescriptor{
		Value:          "INVALIDACCOUNT",
		Message:        "account id  is invalid ",
		Description:    "account is  null",
		HTTPStatusCode: http.StatusBadRequest,
	})
	ErrorCodeIdGetNULL = errcode.Register(errGroup, 20002, errcode.ErrorDescriptor{
		Value:          "NOACCOUNT",
		Message:        "not has account where id is '%s'",
		Description:    "",
		HTTPStatusCode: http.StatusBadRequest,
	})
	ErrorCodeOldPassWdError = errcode.Register(errGroup, 20003, errcode.ErrorDescriptor{
		Value:          "OLDPASSWDERROR",
		Message:        "旧密码错误",
		Description:    "",
		HTTPStatusCode: http.StatusBadRequest,
	})
	ErrorCodeUpdateFAILD = errcode.Register(errGroup, 20004, errcode.ErrorDescriptor{
		Value:          "FAILD",
		Message:        "修改失败",
		Description:    "",
		HTTPStatusCode: http.StatusBadRequest,
	})
	
	
	ErrorCodeRemoveFAILD = errcode.Register(errGroup, 20005, errcode.ErrorDescriptor{
		Value:          "FAILD",
		Message:        "删除失败",
		Description:    "",
		HTTPStatusCode: http.StatusBadRequest,
	})
	
)