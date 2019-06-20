package commons

import (
	md52 "crypto/md5"
	"encoding/hex"
)

func Md5v(str string) string {
	md5 := md52.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum(nil))
}

