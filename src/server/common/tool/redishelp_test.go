package tool

import (
	"fmt"
	"testing"
)

func init() {

}

func Test_Range(t *testing.T) {
	client := GetClient()
	res, err := LRange(client, "a", 0, 100)
	fmt.Println(res)
	fmt.Println(err)

}

func Test_RPushArr(t *testing.T) {
	client := GetClient()
	str := []interface{}{"a", "b"}
	res, err := RPushArr(client, "test", str)
	fmt.Println(res)
	fmt.Println(err)
}
