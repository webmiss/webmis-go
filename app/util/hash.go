package util

import (
	"crypto/md5"
	"encoding/hex"
)

/* MD5 */
func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
