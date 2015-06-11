package types

// KeyID is kind of key or char
type KeyID int

const (
	Begin KeyID = iota
	End
	If
	Then
	While
	Do
	Ret
	Func
	Var
	Const
	Odd
	Write
	WriteLn
	EndOfKeyWd
	Plus
	Minus
	Mult
	Div
	Lparen
	Rparen
	Equal
	Lss
	Gtr
	NotEq
	LssEq
	GtrEq
	Comma
	Period
	Semicolon
	Assign
	EndOfKeySym
	ID
	Num
	Nul
	EndOfToken
	Letter
	Digit
	Colon
	Others
)
