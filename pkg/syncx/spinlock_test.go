// Copyright 2021 Andy Pan & Dietoad. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package syncx

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Benchmark result for three types of locks:
	goos: linux
	goarch: amd64
	pkg: github.com/phamnam2003/ants/pkg/sync
	cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
	BenchmarkMutex-8             	21508472	        55.51 ns/op	       0 B/op	       0 allocs/op
	BenchmarkSpinLock-8          	59300019	        20.37 ns/op	       0 B/op	       0 allocs/op
	BenchmarkBackOffSpinLock-8   	69218594	        17.38 ns/op	       0 B/op	       0 allocs/op
*/

type originSpinLock uint32

func (sl *originSpinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *originSpinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

func NewOriginSpinLock() sync.Locker {
	return new(originSpinLock)
}

func BenchmarkMutex(b *testing.B) {
	m := sync.Mutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Lock()
			//nolint:staticcheck
			m.Unlock()
		}
	})
}

func BenchmarkSpinLock(b *testing.B) {
	spin := NewOriginSpinLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			spin.Lock()
			//nolint:staticcheck
			spin.Unlock()
		}
	})
}

func BenchmarkBackOffSpinLock(b *testing.B) {
	spin := NewSpinLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			spin.Lock()
			//nolint:staticcheck
			spin.Unlock()
		}
	})
}
