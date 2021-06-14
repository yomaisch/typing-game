package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/tjarratt/babble"
)

func main() {
	var (
		ch_rcv = myInput(os.Stdin)
		t      = 1 // タイムリミット
		n      = 0
	)

	babbler := babble.NewBabbler()
	babbler.Count = 1

	fmt.Println("Start the typing game." + " Time limit is " + strconv.Itoa(t) + " minutes." + " Yay, start!")

	for i := true; i; {
		// 英文字をランダムに生成
		q := babbler.Babble()
		fmt.Println(q)

		select {
		case <-time.After(time.Duration(t) * time.Minute): // タイムアウト処理
			fmt.Println("Finished!" + " Your score is " + strconv.Itoa(n) + " points! Good job:)")
			i = false
		case x := <-ch_rcv:
			if x == q {
				fmt.Println("nice")
				n += 1
			} else {
				fmt.Println("incorrect :(")
			}
		}
	}
}

func myInput(r io.Reader) <-chan string {
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch1 <- s.Text()
		}
	}()
	return ch1
}
