package checkup

import (
	"testing"
)

func init() {
	DefaultReceiver = "wenzhengnlin"
}

func TestSend(t *testing.T) {
	Send("xxx", "healthy", "1m")
	Send("xxx", "down", "1m")
	Send("xxx", "healthy", "1m")
	//time.Sleep(3 * time.Second)
}
