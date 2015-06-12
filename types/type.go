package types

import "fmt"

// Max length of name
const MaxName = 31

// Token defines token has fields of KeyId kind and []int32 ([]rune) value
// If kind is key word, second field has []rune value
// If kind is token, second field has []int32 value
type Token struct {
	Kind  KeyID
	ID    [MaxName]byte
	Value int
}

func (t Token) String() string {
	switch {
	case t.Kind == ID:
		return fmt.Sprintf("[%s: %s]", t.Kind, string(t.ID[:MaxName]))
	case t.Kind == Num:
		return fmt.Sprintf("[%s: %d]", t.Kind, t.Value)
	default:
		return fmt.Sprintf("[%s]", t.Kind)
	}
}

// CharClassType has char (byte) key and KeyID value.
var CharClassType = make(map[byte]KeyID)

// InitCharClassType should be called before use CharClassType
func InitCharClassType() {

	for i := byte(0); i < 128; i++ {
		CharClassType[i] = Others
	}
	for i := '0'; i < '9'; i++ {
		CharClassType[byte(i)] = Digit
	}
	for i := 'A'; i < 'Z'; i++ {
		CharClassType[byte(i)] = Letter
	}
	for i := 'a'; i < 'z'; i++ {
		CharClassType[byte(i)] = Letter
	}
	CharClassType['+'] = Plus
	CharClassType['-'] = Minus
	CharClassType['*'] = Mult
	CharClassType['/'] = Div
	CharClassType['('] = Lparen
	CharClassType[')'] = Rparen
	CharClassType['='] = Equal
	CharClassType['<'] = Lss
	CharClassType['>'] = Gtr
	CharClassType['.'] = Period
	CharClassType[','] = Comma
	CharClassType[';'] = Semicolon
	CharClassType[':'] = Colon
}

// KeyWdToResWd has KeyID key and reserved word value.
var KeyWdToResWd = make(map[KeyID]string)

// InitKeyWdToResWd should be called before use KeyWdToResWd
func InitKeyWdToResWd() {
	KeyWdToResWd[Begin] = "begin"
	KeyWdToResWd[End] = "end"
	KeyWdToResWd[If] = "if"
	KeyWdToResWd[Then] = "then"
	KeyWdToResWd[While] = "while"
	KeyWdToResWd[Do] = "do"
	KeyWdToResWd[Ret] = "return"
	KeyWdToResWd[Func] = "function"
	KeyWdToResWd[Var] = "var"
	KeyWdToResWd[Const] = "const"
	KeyWdToResWd[Odd] = "odd"
	KeyWdToResWd[Write] = "write"
	KeyWdToResWd[WriteLn] = "writeln"
	KeyWdToResWd[EndOfKeyWd] = "$dummy1"
	KeyWdToResWd[Plus] = "+"
	KeyWdToResWd[Minus] = "-"
	KeyWdToResWd[Mult] = "*"
	KeyWdToResWd[Div] = "/"
	KeyWdToResWd[Lparen] = "("
	KeyWdToResWd[Rparen] = ")"
	KeyWdToResWd[Equal] = "="
	KeyWdToResWd[Lss] = "<"
	KeyWdToResWd[Gtr] = ">"
	KeyWdToResWd[NotEq] = "<>"
	KeyWdToResWd[LssEq] = "<="
	KeyWdToResWd[GtrEq] = ">="
	KeyWdToResWd[Comma] = ","
	KeyWdToResWd[Period] = "."
	KeyWdToResWd[Semicolon] = ";"
	KeyWdToResWd[Assign] = ":="
	KeyWdToResWd[EndOfKeySym] = "$dummy2"
}
