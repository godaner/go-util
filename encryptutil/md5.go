package encryptutil

import (
	"crypto/md5"
	"fmt"
)

func EncryptMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}