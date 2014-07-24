package main

import (
	"fmt"
	"github.com/doraemon_server/tool/leb128"
	"io"
	"net"
	"time"
)

func test_leb128() {
	bindata := leb128.EncodeULeb128(uint32(300))
	fmt.Println("bindata:", bindata)
}
func main() {

	fmt.Println("server start")
	serverListener, err := net.Listen("tcp", ":9871")
	if err != nil {
		fmt.Println("serverListener create fail!")
	}

	for {

		conn, err := serverListener.Accept()

		if err != nil {
			fmt.Println("accept err")
		} else {
			fmt.Println("accept conn ip:", conn.RemoteAddr())

		}

		go handle_conn(conn)

	}
}

func handle_conn(c net.Conn) {

	fmt.Println("handle_Conn")
	go process_client_read(c)
}

func process_client_read(c net.Conn) {
	defer c.Close()

	//	readBuf := make([]byte, 2048)
	for {
		// Echo all incoming data.
		fmt.Println("ReadStart ")

		io.Copy()
		//readNum, err := c.Read(readBuf)
		//if err != nil {
		//	fmt.Println("client read erro and close")
		//	return
		//} else {
		//	fmt.Println("read ", readBuf[:readNum])
		//}

		err :=

			time.Sleep(time.Second)

	}
}
