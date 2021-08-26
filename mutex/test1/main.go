/**
  @author: dingfeng
  @date: 2021/8/24
  @note:
**/
package main

import (
	"fmt"
	"sync"
	"time"
)
/*
我们看看几种 Mutex 锁的实现:
Barging. 这种模式是为了提高吞吐量，当锁被释放时，它会唤醒第一个等待者，然后把锁给第一个等待者或者给第一个请求锁的人。
Handsoff. 当锁释放时候，锁会一直持有直到第一个等待者准备好获取锁。它降低了吞吐量，因为锁被持有，即使另一个 goroutine 准备获取它。
    一个互斥锁的 handsoff 会完美地平衡两个goroutine 之间的锁分配，但是会降低性能，因为它会迫使第一个 goroutine 等待锁。
Spinning. 自旋在等待队列为空或者应用程序重度使用锁时效果不错。parking 和 unparking goroutines 有不低的性能成本开销，相比自旋来说要慢得多。
*/
func main() {
	done := make(chan struct{}, 1)
	var go1 int
	var go2 int
	var mutex sync.Mutex
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mutex.Lock()
				time.Sleep(time.Microsecond * 100)
				mutex.Unlock()
				go1++
			}
		}
	}()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		time.Sleep(time.Microsecond * 100)
		mutex.Unlock()
		go2++
	}
	done <- struct{}{}
	fmt.Println(go1,go2)
}
