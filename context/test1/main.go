/**
  @author: dingfeng
  @date: 2021/8/24
  @note:
**/

/*
上下文信息传递 （request-scoped），比如处理 http 请求、在请求处理链路上传递信息；
控制子 goroutine 的运行；
超时控制的方法调用；
可以取消的方法调用。
*/
package  main

import (
	"context"
	"time"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			fmt.Println("goroutine exit")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

