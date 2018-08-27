package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	t, _ := tail.TailFile("sample.txt", tail.Config{Follow: true})
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
