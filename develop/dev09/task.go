package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func wget(url string) error {
	// последняя часть урла как название файла
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		log.Fatal("go run task.go <url>")
	}
	err := wget(args[1])
	if err != nil {
		log.Fatal(err)
	}
}
