package main

import (
	"bufio"
	"github.com/vavilen84/codewars/ip_validation/go/validator"
	"log"
	"os"
)

func main() {

	lines := make(chan string)
	quit := make(chan bool)

	inputFile := os.Args[1:][0]
	f, err := os.Open(inputFile)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	go func() {
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		quit <- true
		close(lines)
	}()

	for {
		select {
		case line := <-lines:
			go func(line string) {
				r := validator.IsValidIp(line)
				if r == false {
					log.Println(line)
				}
			}(line)
		case <-quit:
			return
		}
	}
}
