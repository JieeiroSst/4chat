package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"4chat/controller"
)


var (
	conns []net.Conn
)

func MyAccount(a controller.Account) {
	fmt.Println("Name user")
	_, _ = fmt.Scanln(&a.UserName)
	fmt.Println("Password")
	_, _ = fmt.Scanln(&a.Password)
}
func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')

		fmt.Print(msg)
	}
}

func publicMessage(conn net.Conn) {
	for {
		writeMsg := bufio.NewReader(os.Stdin)
		msg, err := writeMsg.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)

		for i := range conns {
			if conns[i] != conn {
				_, _ = conns[i].Write([]byte(msg))
			}
		}
	}
}

func main() {
	var a controller.Account

	connection, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}
	MyAccount(a)
	fmt.Println("********** MESSAGES **********")
	go onMessage(connection)
	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		msg = msg[:len(msg)-1]
		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s-->%s\n", a.UserName, msg)

		_, _ = connection.Write([]byte(msg))

		go publicMessage(connection)
	}

	_ = connection.Close()
}
