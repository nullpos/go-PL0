package main

import (
	"fmt"
	"os"
	"sync"

	"PL0/types"
)

// NextToken output token to ch.
func NextToken(f *os.File, ch chan *types.Token) {
	b := make([]byte, types.MaxName)
	o := 0

	for {
		n, _ := f.ReadAt(b, int64(o))
		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			var token = new(types.Token)
			if b[i] == ' ' || b[i] == '\t' {
				o++
				continue
			}
			cc := types.CharClassType[b[i]]
			switch cc {
			case types.Letter:
				o++
				break
			case types.Digit:
				o++
				break
			case types.Colon:
				if b[i+1] == '=' {
					token.Kind = types.Assign // :=
					o += 2
				} else {
					token.Kind = types.Nul
					o++
				}
				ch <- token
				break
			case types.Lss:
				if b[i+1] == '=' {
					token.Kind = types.LssEq // <=
					o += 2
				} else if b[i+1] == '>' {
					token.Kind = types.NotEq // <>
					o += 2
				} else {
					token.Kind = types.Gtr // <
					o++
				}
				ch <- token
				break
			case types.Gtr:
				if b[i+1] == '=' {
					token.Kind = types.GtrEq // >=
					o += 2
				} else {
					token.Kind = types.Gtr // >
					o++
				}
				ch <- token
				break
			default:
				token.Kind = cc
				o++
				break
			}
		}
	}
	close(ch)
}

func main() {
	types.InitCharClassType()
	filename := "../src/PL0/test.pl1"
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
		NextToken(file, ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

	wg.Wait()
	fmt.Println("Done")
}
