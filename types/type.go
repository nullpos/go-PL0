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

// CharClassType has byte key and KeyID value.
var CharClassType = make(map[byte]KeyID)

// KeyWdToResWd
var KeyWdToResWd = make(map[KeyID]string)

// InitCharClassType should be called before CharClassType
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
