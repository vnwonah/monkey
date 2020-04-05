package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals 
	IDENT = "IDENT" // add, foobar, x, y
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimeters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUMCTION = "FUNCTION"
	LET = "LET"

)

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

