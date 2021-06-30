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
		t      = 1                                          // タイムリミット
		tl     = time.After(time.Duration(t) * time.Minute) // 制限時間処理
		n      = 0
	)

	babbler := babble.NewBabbler()
	babbler.Count = 1

	fmt.Println("Start the typing game." + " Time limit is " + strconv.Itoa(t) + " second." + " Yay, start!")

OuterLoop:
	for {
		// 英文字をランダムに生成
		q := babbler.Babble()
		fmt.Println(q)

		select {
		case <-tl: // 制限時間が来た時の処理
			fmt.Println("Finished!" + " Your score is " + strconv.Itoa(n) + " points! Good job:)")
			break OuterLoop
		case x := <-ch_rcv:
			if x == q {
				fmt.Println("OK!")
				n += 1
			} else {
				fmt.Println("NG :(")
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
