package tool

import (
	//	"fmt"
	//	"gopkg.in/mgo.v2/bson"
	//	"server/common/tool"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"testing"
	//	"io"
	//	"crypto/rand"
	//	"encoding/hex"
	"fmt"
	//	"os"

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

//
//// 3DES加密
//func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
//	block, err := des.NewTripleDESCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	origData = PKCS5Padding(origData, block.BlockSize())
//	// origData = ZeroPadding(origData, block.BlockSize())
//	blockMode := cipher.NewCBCEncrypter(block, key[:8])
//	crypted := make([]byte, len(origData))
//	blockMode.CryptBlocks(crypted, origData)
//	return crypted, nil
//}
//
//// 3DES解密
//func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
//	block, err := des.NewTripleDESCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	blockMode := cipher.NewCBCDecrypter(block, key[:8])
//	origData := make([]byte, len(crypted))
//	// origData := crypted
//	blockMode.CryptBlocks(origData, crypted)
//	origData = PKCS5UnPadding(origData)
//	// origData = ZeroUnPadding(origData)
//	return origData, nil
//}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func testDes() {
	key := []byte("1111111111111111")
	result, err := AesEncrypt([]byte("33@studygolang"), key)
	if err != nil {
		panic(err)
	}
	//	s :=string(result)
	//	 fmt.Sprintf("%x", result)
	//	fmt.Println(s)
	fmt.Println(result)
	bresult := base64.StdEncoding.EncodeToString(result)
	fmt.Println(bresult)
	dres, err := base64.StdEncoding.DecodeString(bresult)
	origData, err := AesDecrypt(dres, key)

	//	origData, err := AesDecrypt([]byte(s), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

//
//func test3Des() {
//	key := []byte("sfe023f_sefiel#fi32lf3e!")
//	result, err := TripleDesEncrypt([]byte("polaris@studygol"), key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(base64.StdEncoding.EncodeToString(result))
//	origData, err := TripleDesDecrypt(result, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(origData))
//}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
func AesEncryptToBase64Str(origData, key []byte) (string, error) {
	result,err:= AesEncrypt(origData, key)
	return base64.StdEncoding.EncodeToString(result),err
}

func AesDecryptByBase64Str(b64crypted string, key []byte) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(b64crypted)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return string(origData), nil
}


func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

//func main() {
//	// DES 加解密
//	testDes()
//	// 3DES加解密
//	test3Des()
//}
func Test_GetAccountManageShopIds(t *testing.T) {
	testDes()
	//	key := []byte("example key 1234")
	//	plaintext := []byte("exampleplaintext")
	//
	//	// CBC mode works on blocks so plaintexts may need to be padded to the
	//	// next whole block. For an example of such padding, see
	//	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	//	// assume that the plaintext is already of the correct length.
	//	if len(plaintext)%aes.BlockSize != 0 {
	//		panic("plaintext is not a multiple of the block size")
	//	}
	//
	//	block, err := aes.NewCipher(key)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	// The IV needs to be unique, but not secure. Therefore it's common to
	//	// include it at the beginning of the ciphertext.
	//	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	//	iv := ciphertext[:aes.BlockSize]
	////	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	////		panic(err)
	////	}
	//
	//	mode := cipher.NewCBCEncrypter(block, iv)
	//	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	//
	//	// It's important to remember that ciphertexts must be authenticated
	//	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	//	// be secure.
	//
	//	fmt.Printf("%x\n", ciphertext)
}

//
//func Test_ExampleNewCFBDecrypter(t *testing.T) {
//	key := []byte("example key 1234")
//	ciphertext, _ := hex.DecodeString("22277966616d9bc47177bd02603d08c9a67d5380d0fe8cf3b44438dff7b9")
//
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//
//	// The IV needs to be unique, but not secure. Therefore it's common to
//	// include it at the beginning of the ciphertext.
//	if len(ciphertext) < aes.BlockSize {
//		panic("ciphertext too short")
//	}
//	iv := ciphertext[:aes.BlockSize]
//	ciphertext = ciphertext[aes.BlockSize:]
//
//	stream := cipher.NewCFBDecrypter(block, iv)
//
//	// XORKeyStream can work in-place if the two arguments are the same.
//	stream.XORKeyStream(ciphertext, ciphertext)
//	fmt.Printf("%s", ciphertext)
//	// Output: some plaintext
//}