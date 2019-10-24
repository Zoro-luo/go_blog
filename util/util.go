package util

import (
	"encoding/base64"
	"github.com/astaxie/beego"
)

//断言 err != nil
func CheckErr(err error,msg interface{}) {
	if err != nil {
		beego.Info(msg, err)
		return
	}
}

// Base64Encoding base64 编码
func Base64Encoding(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

// Base64Decoding base64 解码
func Base64Decoding(enc string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}