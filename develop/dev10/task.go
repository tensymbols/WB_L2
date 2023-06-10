package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout_s := flag.String("timeout", "10s", "timeout")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Fprintf(os.Stderr, "Not enough arguments\n")
		return
	}

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM) //  syscall.SIGQUIT, syscall.SIGTERM
	go Exit(exitCh)

	timeout, err := time.ParseDuration(*timeout_s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Incorrect time\n")
		return
	}

	hostPort := flag.Arg(0) + ":" + flag.Arg(1)

	conn, err := net.DialTimeout("tcp", hostPort, timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connect by %s\n", hostPort)
		return
	}

	defer conn.Close()

	in := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	for {
		fmt.Print("Enter query: ")
		text, err := in.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading query: %v\n", err)
			return
		}
		text = strings.TrimSpace(text)

		fmt.Fprintf(conn, text+"\n")

		if text == "exit" {
			fmt.Println("Close connection")
			return
		}

		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading answer from server: %v\n", err)
			return
		}

		message = strings.TrimSpace(message)
		fmt.Printf("Answer: %s\n", message)
	}

}

func Exit(exitCh chan os.Signal) {
	for {
		select {
		case <-exitCh:
			fmt.Println("Press Ctrl+D\nExit")
			os.Exit(0)
		default:
		}
	}
}
