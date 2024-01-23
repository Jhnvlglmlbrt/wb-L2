package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// execInput обрабатывает ввод пользователя
func execInput(input string, out io.Writer) error {
	cmdArgs := strings.Fields(input)
	if len(cmdArgs) == 0 {
		return nil
	}

	switch cmdArgs[0] {
	case "cd":
		return changeDirectory(cmdArgs)
	case "echo":
		fmt.Fprintln(out, strings.Join(cmdArgs[1:], " "))
	case "ps":
		return executeCommand(exec.Command("ps", "aux"), out)
	case "pwd":
		return printWorkingDirectory(out)
	case "kill":
		return killProcess(cmdArgs)
	case "q":
		os.Exit(0)
	}

	return nil
}

// changeDirectory обрабатывает команду cd
func changeDirectory(cmdArgs []string) error {
	if len(cmdArgs) < 2 {
		return os.Chdir(os.Getenv("HOME"))
	}
	return os.Chdir(cmdArgs[1])
}

// executeCommand выполняет переданную команду
func executeCommand(cmd *exec.Cmd, out io.Writer) error {
	cmd.Stderr = os.Stderr
	cmd.Stdout = out
	return cmd.Run()
}

// printWorkingDirectory выводит текущую рабочую директорию
func printWorkingDirectory(out io.Writer) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(out, dir)
	return nil
}

// killProcess обрабатывает команду kill
func killProcess(cmdArgs []string) error {
	if len(cmdArgs) < 2 {
		return fmt.Errorf("usage: kill <PID>")
	}

	pid := cmdArgs[1]
	cmd := exec.Command("taskkill", "/F", "/PID", pid)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

// runCommand выполняет команду с входными данными
func runCommand(cmd *exec.Cmd, input []byte) ([]byte, error) {
	cmd.Stdin = bytes.NewReader(input) // *bytes.Reader используя массив байт input устанвливается в качестве ввода Stdin
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return output.Bytes(), err
}

// getCommands возвращает массив команд из строки пайплайна
func getCommands(input string) []*exec.Cmd {
	commands := make([]*exec.Cmd, 0)

	strCommands := strings.Split(input, "|")
	for _, cmd := range strCommands {
		cmdArgs := strings.Fields(cmd)
		commands = append(commands, exec.Command(cmdArgs[0], cmdArgs[1:]...))
	}

	return commands
}

// Pipeline - итерируемся по командам и выполняем, записывая результат в output, печатаем в os.Stdout output и
// данные выполнения сохраняем в inputBuffer, и затем запускаем следующую команду с входными данными прошлой и т.д.
func Pipeline(input string) error {
	commands := getCommands(input)
	if len(commands) < 1 {
		return nil
	}

	// служит для передачи вывода одной команды в качестве входных данных для следующей команды в пайплайне
	var inputBuffer []byte
	for _, cmd := range commands {
		output, err := runCommand(cmd, inputBuffer)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(output))
		inputBuffer = output // результат выполнения команды записывается в буфер и будет использоваться как вход для след команды
	}

	return nil
}

// // Fork - функция обработки форк команд в дочернем процессе запускаем переданную команду
// func Fork(input string) error {
// 	input = strings.TrimSuffix(input, " &")
// 	fmt.Println(input)
// 	id, _, errno := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
// 	if errno != 0 {
// 		os.Exit(1)
// 	}
// 	if id == 0 {
// 		err := execInput(input, nil)
// 		if err != nil {
// 			return err
// 		}
// 		os.Exit(0)
// 	}
// 	return nil
// }

// readInput читает ввод пользователя
func readInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Printf("%s> ", dir)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		input = strings.TrimSuffix(input, "\n")

		if strings.Contains(input, "|") {
			err := Pipeline(input)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		if err := execInput(input, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func main() {
	readInput()
}
