package main

import (
	"flag"
	"fmt"
	"os"
)

// Config 파라미터 처리
type Config struct {
	pid uint // 감시할 process의 pid
}

// Check conf 체크
func (p *Config) Check() {
	if p.pid == 0 {
		fmt.Println("Invaild PID!")
		usage()
		os.Exit(0)
	}
}

// Print Info 로그로 출력한다.
func (p *Config) Print() {
	fmt.Printf("PID: %d\n", p.pid)
}

func usage() {
	fmt.Printf("Usage: %s arguement...\n", os.Args[0])
	flag.PrintDefaults()
}
