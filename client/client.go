package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Account struct {
	UserName string
	Password string
}

var (
	conns []net.Conn
)

func MyAccount(a Account) {
	fmt.Println("Name user")
	fmt.Scanln(&a.UserName)
	fmt.Println("Password")
	fmt.Scanln(&a.Password)
}

func Accecpt() {
	var status string

	fmt.Println("can you speack statusd")
	fmt.Println("yes/no")
	fmt.Scan(&status)
	switch status {
	case "yes":
	case "no":
		break
		return
	}
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
				conns[i].Write([]byte(msg))
			}
		}
	}
}

func main() {
	var a Account

	connection, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}

	MyAccount(a)
	Accecpt()
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')

	nameInput = nameInput[:len(nameInput)-1]

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

		connection.Write([]byte(msg))

		go publicMessage(connection)
	}

	connection.Close()
}
