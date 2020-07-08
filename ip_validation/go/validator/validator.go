package validator

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsValidIp(i string) (r bool) {
	pattern := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	r, err := regexp.Match(pattern, []byte(i))
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func IsValidIpV2(i string) (r bool) {
	digitsCount := 4
	splitted := strings.Split(i, ".")
	if len(splitted) != digitsCount {
		return
	}
	for _, v := range splitted {
		var b []byte
		b = []byte(v)
		if string(b[0]) == `0` {
			return
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			println(err)
			os.Exit(1)
		}
		if i > math.MaxUint8 || i < 1 {
			return false
		}
	}
	return true
}
