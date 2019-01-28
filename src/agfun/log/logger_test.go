package log

import (
	"testing"
	"time"
)

func TestInitLog(t *testing.T) {
	InitLog("myTest", ".", 3, 3*time.Hour, 10*time.Second)
	Println("test,test")
	tick := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-tick.C:
			Println("111111111111111111111")
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
