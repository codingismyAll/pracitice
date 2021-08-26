/**
  @author: dingfeng
  @date: 2021/8/24
  @note:
**/
package main

import (
	"../test1/faith"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
)

func main()  {
	//testTcpConnPool()
	con:=testType()
	if pc, ok := con.(*pool.PoolConn); ok {
		pc.MarkUnusable()
		pc.Close()
	}else{
		fmt.Println(false)
		dir,_:=os.Getwd()
		fmt.Println(dir)
	}
}
func testTcpConnPool() {
	sp2 := sync.Pool{
		New: func() interface{} {
			conn, err := net.Dial("tcp", "127.0.0.1:9277")
			if err != nil {
				return nil
			}
			return conn
		},
	}
	buf := make([]byte, 1024)
	//获取对象
	conn := sp2.Get().(net.Conn)
	//使用对象
	conn.Write([]byte("GET / HTTP/1.1 \r\n\r\n"))
	n, _ := conn.Read(buf)
	fmt.Println("conn read : ", string(buf[:n]))
	//打印conn的地址
	fmt.Println(conn)
	//把对象放回池中
	sp2.Put(conn)
	//我们人为的进行一次垃圾回收
	runtime.GC()
	//再次获取池中的对象
	conn2 := sp2.Get().(net.Conn)
	//这时发现conn2的地址与上面的conn的地址不一样了
	//说明池中我们之前放回的对象被全部清除了，显然这并不是我们想看到的
	//所以sync.Pool不适合用于scoket长连接或数据库连接池
	fmt.Println(conn2)
}

func testType() net.Conn{
	conn, err := net.Dial("tcp", "127.0.0.1:9277")
	if err != nil {
		return nil
	}
	return conn
}