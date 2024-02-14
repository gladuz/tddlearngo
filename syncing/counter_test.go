package syncing

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("running 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		wanted := 3
		assertCounterEqual(t, &counter, wanted)
	})

	t.Run("running in goroutine", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounterEqual(t, &counter, wantedCount)
	})
}

func assertCounterEqual(t *testing.T, counter *Counter, wanted int) {
	if counter.Value() != wanted {
		t.Errorf("got %d, wanted %d", (*counter).Value(), wanted)
	}
}
