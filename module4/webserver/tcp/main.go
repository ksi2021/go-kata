package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Read the request
	request, _ := reader.ReadString('\n')
	fmt.Println(request)

	// Parse the request
	parts := strings.Split(request, " ")
	if len(parts) < 2 {
		return
	}
	method := parts[0]
	path := parts[1]

	// Write the response
	if method == "GET" && path == "/" {
		m := "HTTP/1.1 200 OK\n"
		m = m + "Server: nginx/1.19.2\n"
		m = m + "Date: Mon, 19 Oct 2020 13:13:29 GMT\n"
		m = m + "Content-Type: text/html\n"
		m = m + "Content-Length: 98\n" // указываем размер контента
		m = m + "Last-Modified: Mon, 19 Oct 2020 13:13:13 GMT\n"
		m = m + "Connection: keep-alive\n"
		m = m + "ETag: \"5f8d90e9-62\"\n"
		m = m + "Accept-Ranges: bytes\n"
		m = m + "\n"
		m = m + "<!DOCTYPE html>\n"
		m = m + "<html>\n"
		m = m + "<head>\n"
		m = m + "<title>Webserver</title>\n"
		m = m + "</head>\n"
		m = m + "<body>\n"
		m = m + "hello world\n"
		m = m + "</body>\n"
		m = m + "</html>\n"
		conn.Write([]byte(m))
	} else if method == "POST" && path == "/upload" {
		// Read the file from the request body
		file, _ := os.Create("uploaded_file.txt")
		defer file.Close()
		//io.Copy(file, reader)
		//fmt.Println(conn)
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\nFile uploaded successfully!")
		return
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 403 Not Found\r\n\r\n")
	}
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}
