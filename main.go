package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"time"
)

var CRLF = "\r\n"

func handleConnection(conn net.Conn, keepAlive int) {
	defer conn.Close()
	for keepAlive > 0 {
		keepAlive--
		b := make([]byte, 1000)
		for count, _ := conn.Read(b); count == 0; {
			time.Sleep(1)
		}
		body := "<h1>Hello world</h1>"
		header := "HTTP/1.1 200 OK" + CRLF +
			"Date: " + time.Now().Format(time.RFC1123) + CRLF +
			"Content-Length: " + strconv.Itoa(len(body)) + CRLF +
			"Content-Type: text/html;charset=utf-8" + CRLF
		if keepAlive > 0 {
			header += "Connection: keep-alive" + CRLF
		} else {
			header += "Connection: close" + CRLF
		}
		conn.Write([]byte(header + CRLF + body))
		// fmt.Println(b)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	ln, err := net.Listen("tcp", ":8080")
	if err == nil {
		for {
			conn, err := ln.Accept()
			if err == nil {
				go handleConnection(conn, 500)
			}
		}
	}
	fmt.Println("Error!")
}
