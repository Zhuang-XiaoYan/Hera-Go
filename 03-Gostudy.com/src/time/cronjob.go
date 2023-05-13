package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	n := 5
	for t := range ticker.C {
		n--
		fmt.Println(t)
		if n == 0 {
			ticker.Stop()
			break
		}
	}
}
