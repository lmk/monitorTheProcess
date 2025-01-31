package main

import (
	"flag"
	"io/ioutil"
	"os"
	"time"

	"github.com/struCoder/pidusage"
)

var conf Config

// cpu 사용량 계산 기준
// mem 사용량
// ps -o pcpu,rss -p 4179 와 같음.
// rss * 1024

func main() {

	// 로거 초기화
	InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Info.Println("Start App")

	flag.Parse()
	conf.Check()
	conf.Print()

	Info.Printf("ps -o pcpu,rss -p %d", conf.pid)

	var sumInfo pidusage.SysInfo
	showTime := time.Now().Truncate(time.Second).Add(time.Second * time.Duration(conf.duration))

	for {

		nextTime := time.Now().Truncate(time.Second).Add(time.Second)

		fsysInfo, err := pidusage.GetStat(int(conf.pid))
		if err != nil {
			panic(err)
		}

		sumInfo.CPU += fsysInfo.CPU
		sumInfo.Memory += fsysInfo.Memory / 1024

		if showTime.Before(nextTime) {
			Info.Printf("pcpu,rss: %.2f %.0f\n", sumInfo.CPU/float64(conf.duration), sumInfo.Memory/float64(conf.duration))
			sumInfo = pidusage.SysInfo{}
			showTime = showTime.Add(time.Second * time.Duration(conf.duration))
		}
		//Info.Printf("pcpu,rss: %.2f %.0f\n", fsysInfo.CPU, fsysInfo.Memory)
		time.Sleep(time.Until(nextTime))
	}
}

func init() {
	flag.UintVar(&conf.pid, "pid", 0, "PID")
	flag.Uint64Var(&conf.duration, "duration", 60, "Duration Second")
}
