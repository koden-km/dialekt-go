package dialekt

import (
	"strings"
	// 	"regex"
)

type State int

const (
	_ State = iota
	StateBegin
	StateSimpleString
	StateQuotedString
	StateQuotedStringEscape
)

type Lexer struct {
	currentOffset int
	currentLine   int
	currentColumn int
	state         State
	tokens        []*Token
	nextToken     *Token
	buffer        string
}

func NewLexer() *Lexer {
	lexer := &Lexer{
		0,
		1,
		0,
		StateBegin,
		make([]*Token),
		nil,
		"",
	}

	return lexer
}

// The expression to tokenize.
// Returns the tokens of the expression or an error if the expression is invalid.
func (lexer *Lexer) Lex(expression string) (tokens []Token, err *ParseError)
