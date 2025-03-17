package stream

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
)

type FileServer struct{}

const PORT = "127.0.0.1:3000"

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n == 0 {
			return
		}

		file := buf[:n]
		fmt.Println(string(file))
	}
}

func sendFile(size int) error {
	file := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", PORT)
	if err != nil {
		return err
	}

	n, err := conn.Write(file)
	if err != nil {
		return err
	}

	fmt.Printf("written %d bytes over the network\n", n)

	return nil
}
