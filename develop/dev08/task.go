package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	buf := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		cmd, err := buf.ReadString('\n')
		cmd = strings.TrimSuffix(cmd, "\n")
		cmd = strings.TrimSpace(cmd)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = command(cmd); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func command(cmd string) error {
	cmdArg := strings.Split(cmd, " ")
	switch cmdArg[0] {
	case "cd":
		_, err := exec.Command("bash", "-c", "cd", cmdArg[1]).Output()
		if err != nil {
			return err
		}
	case "echo":
		out, err := exec.Command("echo", strings.Join(cmdArg[1:], " ")).Output()
		if err != nil {
			return err
		}

		fmt.Printf("%s", out)
	case "pwd":
		out, err := exec.Command("bash", "-c", "pwd").Output()
		if err != nil {
			return err
		}

		fmt.Printf("%s", out)
	case "kill":
		out, err := exec.Command("kill", cmdArg[1]).Output()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", out)
	case "ps":
		out, err := exec.Command("bash", "-c", "ps").Output()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", out)
	case "quit":
		os.Exit(0)
	}
	return nil
}
