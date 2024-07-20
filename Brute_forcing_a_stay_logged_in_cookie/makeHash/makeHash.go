package makehash

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func MakeHash(username string, password string) string {

	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	stringTobase64 := username + ":" + password
	return base64.RawStdEncoding.EncodeToString([]byte(stringTobase64))

}
