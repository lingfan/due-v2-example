package xmd5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5(str string) string {
	h := md5.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
