package semaphore_test

import (
	"testing"

	"github.com/kamilsk/semaphore"
)

// TODO
// - [ ] investigate benchmarks from github.com/marusama/semaphore
// - [ ] add benchmarks for github.com/dropbox/godropbox
// - [ ] test overhead of full semaphore

func BenchmarkSemaphore_Acquire(b *testing.B) {
	sem := semaphore.New(b.N)

	for i := 0; i < b.N; i++ {
		_, _ = sem.Acquire(nil)
	}

	if sem.Occupied() != sem.Capacity() {
		b.Error("full filled semaphore is expected")
	}
}

func BenchmarkSemaphore_Acquire_Release(b *testing.B) {
	sem := semaphore.New(b.N)

	for i := 0; i < b.N; i++ {
		_, _ = sem.Acquire(nil)
		_ = sem.Release()
	}

	if sem.Occupied() != 0 {
		b.Error("empty semaphore is expected")
	}
}
