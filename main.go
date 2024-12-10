package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {

	// Код для измерения

	for {

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("kwdb-send: ")
		message, _ := reader.ReadString('\n')
		start := time.Now()
		for i := 0; i < 20; i++ {
			go goe(i)
		}
		go send(message)

		duration := time.Since(start)
		// например, "2h3m0.5s" или "4.503μs"
		fmt.Printf("time: %v  , nano: %v", duration.Seconds(), duration.Nanoseconds())
	}

}

func goe(pac int) {

	for i := pac * 1000; i < (pac*1000)+1000; i++ {
		send("SET value=cacheafgljgfjkgfjklgfdsjkgdfkjlgdfsljkgfdsljkgfdljkfgdsljkgfdsljkgfdsljk" + strconv.Itoa(i) + " key=" + strconv.Itoa(i))
	}
}
func send(message string) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// отправляем сообщение серверу
	if n, err := conn.Write([]byte(message)); n == 0 || err != nil {
		fmt.Println(err)
		return
	}
	// получем ответ

	buff := make([]byte, 1024)
	_, err = conn.Read(buff)
	if err != nil {
		return
	}

	response := string(bytes.TrimRight(buff, "\x00, \n"))

	fmt.Println("Ответ:", response)
	conn.Close()
}
