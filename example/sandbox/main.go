package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

const factor = 60

func main() {
	fmt.Printf("time.Now(): '%#v'\n", time.Now())
	fmt.Printf("time.Now().UnixNano(): '%#v'\n", time.Now().UnixNano())
	fmt.Printf("time.Now().UnixNano()/factor: '%#v'\n", time.Now().UnixNano()/factor)
	fmt.Println()
	fmt.Printf("time.Now().UnixNano()/1e9: '%#v'\n", time.Now().UnixNano()/1e9)
	fmt.Printf("time.Now().UnixNano()/1e9/factor: '%#v'\n", time.Now().UnixNano()/1e9/factor)
	fmt.Println()

	m := map[int]int{}
	fmt.Printf("m: %#v\n", m)
	m[42]++
	fmt.Printf("m: %#v\n", m)
	fmt.Println()

	fmt.Printf("levels: %#v\n", log.AllLevels)
	fmt.Printf("debug: %#v\n", log.DebugLevel)
	fmt.Printf("trace: %#v\n", log.TraceLevel)
}
