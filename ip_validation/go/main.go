package _go

import (
	"fmt"
	"regexp"
)

func IsValidIp(i string) (r bool) {
	pattern := `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	r, err := regexp.Match(pattern, []byte(i))
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
