package encryptutil

import (
	"crypto/md5"
	"strings"
	"encoding/hex"
)

func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}
func EncodeMd5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}