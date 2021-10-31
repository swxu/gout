package util

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
)

type Func func(ctx context.Context)

func GoBatchFnWithFixedGoroutines(ctx context.Context, maxGoroutines int, fns []Func) {
	if maxGoroutines <= 0 {
		panic("[GoBatchFnWithFixedGoroutines] `maxGoroutines` must be greater than 0")
	}
	guard := make(chan struct{}, maxGoroutines)
	defer func() {
		go func() {
			close(guard)
		}()
	}()
	wg := &sync.WaitGroup{}
	for _, fn := range fns {
		guard <- struct{}{}
		wg.Add(1)
		go RunFnWithGuard(ctx, wg, fn, guard)
	}
	wg.Wait()
}

func RunFnWithGuard(ctx context.Context, wg *sync.WaitGroup, fn Func, guard chan struct{}) {
	defer func() {
		<-guard
	}()
	RunFn(ctx, wg, fn)
}

func GoBatchFn(ctx context.Context, fns []Func) {
	wg := &sync.WaitGroup{}
	for _, fn := range fns {
		wg.Add(1)
		go RunFn(ctx, wg, fn)
	}
	wg.Wait()
}

func RunFn(ctx context.Context, wg *sync.WaitGroup, fn Func) {
	defer func() {
		wg.Done()
		RoutineRecovery(ctx)
	}()
	fn(ctx)
}

func RoutineRecovery(ctx context.Context) {
	if e := recover(); e != nil {
		fmt.Printf("panic stack:%+v, err=%+v", string(debug.Stack()), e)
	}
}
