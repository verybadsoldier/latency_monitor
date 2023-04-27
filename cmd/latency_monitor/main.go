package main

import (
	"fmt"
	"os"
	"time"
)

func openFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func main() {
	var interval int64 = 1000

	fout, err := openFile("latency.log")
	if err != nil {
		panic(err)
	}

	ferr, err := openFile("latency_err.log")
	if err != nil {
		panic(err)
	}

	var prev_now int64
	for {
		nowTime := time.Now()
		now := time.Now().UnixMilli()
		if prev_now != 0 {
			diff := now - prev_now - interval

			msg := fmt.Sprintf("%s - %d", nowTime.Format(time.RFC3339), diff)
			fmsg := fmt.Sprintf("%s\n", msg)
			fout.WriteString(fmsg)
			fmt.Println(msg)
			if diff > 500 {
				ferr.WriteString(fmsg)
			}
		}
		time.Sleep(time.Duration(interval * 1000 * 1000))

		prev_now = now
	}
}
