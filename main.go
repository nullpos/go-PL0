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
			for i, l := 0, len(b); i < l; i++ {
				token := new(types.Token)
				if b[i] == ' ' || b[i] == '\t' {
					continue
				}
				func(c byte) {
					cc := types.CharClassType[c]
					switch cc {
					case types.Letter:
						tmp := make([]byte, types.MaxName)
						itr := 0
						for {
							tmp[itr] = c
							i++
							itr++
							if i == l {
								i--
								break
							}
							c = b[i]
							cc = types.CharClassType[c]
							if cc != types.Digit && cc != types.Letter {
								i--
								break
							}
						}
						if itr >= types.MaxName {
							fmt.Println("too long")
						}

						for j := types.KeyID(0); j < types.EndOfKeyWd; j++ {
							if types.KeyWdToResWd[j] == string(tmp[:itr]) {
								token.Kind = j
								return
							}
						}
						token.Kind = types.ID
						token.ID = string(tmp[:])
						return
					case types.Digit:
						num := 0
						for {
							num = 10*num + int(c-'0')
							i++
							if i == l {
								i--
								break
							}
							c = b[i]
							cc = types.CharClassType[c]
							if cc != types.Digit {
								i--
								break
							}
						}
						if num>>31 > 0 {
							fmt.Println("too large")
						}
						token.Kind = types.Num
						token.Value = num
						return
					case types.Colon:
						if i+1 < l {
							n := b[i+1]
							switch {
							case n == '=':
								token.Kind = types.Assign // :=
								i++
								return
							default:
								token.Kind = types.Nul
								return
							}
						} else {
							token.Kind = cc
							return
						}
					case types.Lss:
						if i+1 < l {
							n := b[i+1]
							switch {
							case n == '=':
								token.Kind = types.LssEq // <=
								i++
								return
							case n == '>':
								token.Kind = types.NotEq // <>
								i++
								return
							default:
								token.Kind = types.Lss // <
								return
							}
						} else {
							token.Kind = cc
							return
						}
					case types.Gtr:
						if i+1 < l {
							n := b[i+1]
							switch {
							case n == '=':
								token.Kind = types.GtrEq // >=
								i++
								return
							default:
								token.Kind = types.Gtr // >
								return
							}
						} else {
							token.Kind = cc
							return
						}
					default:
						token.Kind = cc
						return
					}
				}(b[i])
				ch <- token
			}
		} else {
			break
		}
	}
	close(ch)
}

func main() {
	types.InitCharClassType()
	types.InitKeyWdToResWd()
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

	//*
	ofilename := "../go-PL0/test-convert.pl1"
	ofile, oe := os.OpenFile(ofilename, os.O_WRONLY|os.O_CREATE, 0600)
	if oe != nil {
		fmt.Println(oe)
		return
	}
	defer file.Close()
	for i := range ch {
		ofile.WriteString(i.String())
		ofile.WriteString("\n")
	}
	/*/
	for i := range ch {
		fmt.Println(i)
	}
	//*/

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
