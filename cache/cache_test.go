package cache

import "testing"

func TestCache(t *testing.T) {
	user, content, status, expire := "u", "body", "ok", "1m"
	Set(user, content, status, expire)
	_, found := Get(user, content, status)
	if !found {
		t.Errorf("cache should be exist")
	}
	_, found = Get(user, content, "nook")
	if found {
		t.Errorf("cache should not exist")
	}
}
