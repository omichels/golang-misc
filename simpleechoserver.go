// $ nc localhost 3540

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"text/template"
)

const PORT = 3540

func main() {
	server, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if server == nil {
		panic("couldn't start listening: " + err.Error())
	}
	conns := clientConns(server)
	for {
		go handleConn(<-conns)
	}
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Printf("couldn't accept: " + err.Error())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

type Answer struct {
	In  string
	Out string
}

func (a *Answer) processOut() {
	if strings.Contains(a.In, "foo") {
		a.Out = "Har, harrrr"
	} else {
		a.Out = a.In
	}
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		answerData := Answer{
			In: string(line[:len(line)-1]),
		}
		answerData.processOut()
		t, err := template.New("ans").Parse("You said \"{{ .In}}\", but we say: \"{{ .Out}}\"\n")
		if err != nil {
			panic(err)
		}
		buf := &bytes.Buffer{}
		err = t.Execute(buf, answerData)

		client.Write(buf.Bytes())
	}
}
