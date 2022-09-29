package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// 生成一定长度的密码
func GetRandomChar(pl int64, charset string) (res string) {
	charsetLen := len(charset)
	for i := int64(0); i < pl; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(charsetLen)))
		char := string(charset[n.Int64()])
		res += char
	}
	//替换第一个为小写字母
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(26)))
	char := string(charset[n.Int64()])
	res = strings.Replace(res, string(res[0]), char, 1)
	return res
}
