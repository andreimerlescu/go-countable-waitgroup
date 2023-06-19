package countable_waitgroup

import (
	"testing"
	"time"
)

func TestCountableWaitGroup(t *testing.T) {
	var wg CountableWaitGroup

	if wg.Count() != 0 || wg.IsPending() {
		t.Errorf("Initial count or pending status incorrect")
	}

	wg.Add(1)
	if wg.Count() != 1 || !wg.IsPending() {
		t.Errorf("Count or pending status incorrect after Add")
	}

	wg.Done()
	if wg.Count() != 0 || wg.IsPending() {
		t.Errorf("Count or pending status incorrect after Done")
	}

	wg.Add(1)
	wg.PreventAdd()
	if wg.CanAdd() {
		t.Errorf("CanAdd failed to indicate whether further adds are permitted")
	}
	wg.Add(1)
	if wg.Count() != 1 || !wg.IsPending() {
		t.Errorf("PreventAdd failed to prevent further adds")
	}
}

func TestCountableWaitGroupWithConcurrency(t *testing.T) {
	var wg CountableWaitGroup
	var workers = 10
	var jobsPerWorker = 1000

	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < jobsPerWorker; j++ {
				wg.Add(1)
				go func() {
					time.Sleep(1 * time.Millisecond)
					wg.Done()
				}()
			}
		}()
	}

	wg.Wait()
	if wg.Count() != 0 || wg.IsPending() {
		t.Errorf("Count or pending status incorrect after all jobs are done")
	}
}
