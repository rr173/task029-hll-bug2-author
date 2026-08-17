package hll

import (
	"fmt"
	"sync"
	"testing"
)

func TestProbeConcurrentAddAndEstimate(t *testing.T) {
	h, err := New(10)
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for worker := 0; worker < 8; worker++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				h.Add([]byte(fmt.Sprintf("%d-%d", worker, i)))
				_ = h.Estimate()
			}
		}(worker)
	}
	wg.Wait()
}
