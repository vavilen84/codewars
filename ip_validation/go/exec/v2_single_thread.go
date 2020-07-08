package main

import (
	"bufio"
	"fmt"
	"github.com/vavilen84/codewars/ip_validation/go/validator"
	"log"
	"os"
)

func main() {
	inputFile := os.Args[1:][0]
	f, err := os.Open(inputFile)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ip := scanner.Text()
		r := validator.IsValidIpV2(ip)
		if r == false {
			fmt.Println(ip)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
