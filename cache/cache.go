// cache to avoid repeat send notice message
package cache

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var C = cache.New(10*time.Minute, 10*time.Minute)

type msg struct {
	content string
	status  string
	d       time.Duration
}

// how to detect it's new ( by parameter )
// have cache ( status change is a new )
// compare the old status and new status
func Set(user, content, status, expire string) {
	d, _ := time.ParseDuration(expire)
	C.Set(user+content, msg{content, status, d}, d)
}

func Get(user, content, status string) (time.Time, bool) {
	cc, d, found := C.GetWithExpiration(user + content)
	if found {
		if cmsg, ok := cc.(msg); ok && cmsg.content == content {
			if cmsg.status == status {
				// cache exist and status not changed, return true
				return d, true
			} else {
				// if status changed, update the cache
				C.Set(user+content, msg{content, status, cmsg.d}, cmsg.d)
			}
		}
	}
	return time.Time{}, false
}
