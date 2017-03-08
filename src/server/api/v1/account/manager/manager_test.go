package manager

import (
	"fmt"
	"server/common/tool"
	"testing"
)

func Test_GetAccountsPage(t *testing.T) {
	res, err := GetAccountsPage(tool.Meta{Limit: 100}, nil)
	fmt.Println(res)
	fmt.Println(err)
}
