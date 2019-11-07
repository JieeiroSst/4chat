package main

import (
	"4chat/db"
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	conns   []net.Conn
	connch  = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgch   = make(chan string)
)

type Account struct {
	UserName string
	Password string
}

func manager(a Account) {

}

func main() {
	server, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	go db.HandleDataBase()

	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatal(err)
			}

			conns = append(conns, conn)
			connch <- conn
		}
	}()

	for {
		select {
		case conn := <-connch:
			go onMessage(conn)
		case msg := <-msgch:
			go fmt.Print(msg)
		case conn := <-closeCh:
			fmt.Println("client exit")
			removeConn(conn)
		}
	}
}

func removeConn(conn net.Conn) {
	var i int
	for i = range conns {
		if conns[i] == conn {
			break
		}
	}
	conns = append(conns[i:], conns[:i+1]...)
}

func publicMsg(conn net.Conn, msg string) {
	for i := range conns {
		if conns[i] != conn {
			conns[i].Write([]byte(msg))
		}
	}
}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msgch <- msg
		publicMsg(conn, msg)
	}

	closeCh <- conn
}
