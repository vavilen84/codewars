package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	digitsCount = 4
	filename    = "address.txt"
)

func main() {
	count, err := strconv.Atoi(os.Args[1:][0])
	if err != nil {
		println(err)
		os.Exit(1)
	}
	f, err := os.Create(filename)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer f.Close()
	generateRandIps(count, f)

}

func generateRandIps(count int, f *os.File) {
	for i := 0; i < count; i++ {
		ip := generateRandIp()
		_, err := f.Write([]byte(ip))
		if err != nil {
			println(err)
			os.Exit(1)
		}
	}
}

func generateRandIp() string {
	ip := ""
	for i := 0; i < digitsCount; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		ip += strconv.Itoa(r.Intn(500))
		if i <= digitsCount-2 {
			ip += "."
		}
	}
	return ip + "\n"
}
