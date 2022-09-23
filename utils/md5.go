package utils
//黑名单相关逻辑
import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 token 编码
func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}