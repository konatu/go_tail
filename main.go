package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename := ""
	tails, err := tail.TailFile(filename, tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false, //filename にファイルが存在していなかったらエラーとする
		Poll:      true,
		Follow:    true,
	})
	if err != nil {
		fmt.Println("tail file error:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
