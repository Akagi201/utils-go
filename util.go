package util

import (
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
