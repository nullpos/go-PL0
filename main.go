package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"go-PL0/types"
)

var b []byte

// GetToken output token to ch.
func GetToken(f *os.File, ch chan *types.Token) {
	getNextLine := nextLiner(f)
	for {
		if getNextLine() {

		} else {
			break
		}
	}
	close(ch)
}

func main() {
	types.InitCharClassType()
	filename := "../go-PL0/test.pl1"
	file, e := os.Open(filename)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup

	ch := make(chan *types.Token)
	wg.Add(1)

	go func() {
		defer wg.Done()
		GetToken(file, ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

	wg.Wait()
	fmt.Println("Done")
}

func nextLiner(f *os.File) func() bool {
	scanner := bufio.NewScanner(f)
	return func() bool {
		if scanner.Scan() {
			b = scanner.Bytes()
			return true
		}
		return false
	}
}
