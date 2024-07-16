package common_test

import (
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	t.Logf("before: %v", time.Now().String())
	timer1 := time.NewTimer(10 * time.Second)
	done := make(chan bool)
	go func() {
		<-timer1.C
		done <- true
	}()

	<-done

	t.Logf("after: %v", time.Now().String())
}
