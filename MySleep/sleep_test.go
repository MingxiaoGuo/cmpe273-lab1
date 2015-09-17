package main
import "testing"
import "time"
func Test123(t *testing.T){
	now := time.Now()
	mySleep(1)
	end := time.Since(now)
	if end < 1 {
		t.Error("Expect 1, got", end)
	}
}
