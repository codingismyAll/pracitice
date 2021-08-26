/**
  @author: dingfeng
  @date: 2021/8/24
  @note:
**/
package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"time"
)

/*
并行工作流
错误处理 或者 优雅降级
context 传播和取消
利用局部变量+闭包
https://pkg.go.dev/golang.org/x/sync/errgroup
*/
func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		time.Sleep(time.Second * 5)
		return errors.New("test err")
	})
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("tuichu")
				return nil
			default:
				fmt.Println("do sth")
				time.Sleep(time.Second * 1)
			}
		}
	})
	if err := g.Wait(); err != nil {
		fmt.Println(ctx.Err())
		fmt.Println(err)
	}

}
