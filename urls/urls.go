// Package urls contains some url related helper functions.
package urls

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// ParseRawStreamURL parse RTMP url
func ParseRawStreamURL(rawStreamURL string) (vhost, app, stream string) {
	s := strings.Split(rawStreamURL, "/")
	vhost = s[0]
	app = s[1]
	stream = strings.Join(s[2:], "/")
	return vhost, app, stream
}

// GetMD5Hash calc a string md5
func GetMD5Hash(text string) string {
	hasher := md5.New()
	_, err := hasher.Write([]byte(text))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
