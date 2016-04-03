package dialekt

import (
	"strings"
	"unicode"
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
func (lexer *Lexer) Lex(expression string) (tokens []Token, err *ParseError) {
	lexer.currentOffset = 0
	lexer.currentLine = 1
	lexer.currentColumn = 0
	lexer.state = StateStateBegin
	lexer.tokens = make([]*Token)
	lexer.buffer = ""

	length := len(expression)
	previousChar := ' '

	for lexer.currentOffset < length {
		char = expression[lexer.currentOffset]
		lexer.currentColumn++

		if '\n' == previousChar || ('\r' == previousChar && '\n' != char) {
			lexer.currentLine++
			lexer.currentColumn = 1
		}

		if StateSimple_STRING == lexer.state {
			lexer.handleSimpleStringState(char)
		} else if StateQuotedString == lexer.state {
			lexer.handleQuotedStringState(char)
		} else if StateQuotedStringEscape == lexer.state {
			lexer.handleQuotedStringEscapeState(char)
		} else {
			lexer.handleBeginState(char)
		}

		lexer.currentOffset++
		previousChar = char
	}

	if StateSimpleString == lexer.state {
		lexer.finalizeSimpleString()
	} else if StateQuotedString == lexer.state {
		return nil, NewParseError("Expected closing quote.")
	} else if StateQuotedStringEscape == lexer.state {
		return nil, NewParseError("Expected character after backslash.")
	}

	return lexer.tokens, nil
}

func (lexer *Lexer) handleBeginState(char rune) {
	// TODO: look up what PHP says this checks.
	// if (ctype_space($char)) {
	// if char == ' ' {
	if unicode.IsSpace(char) {
		// ignore ...
	} else if char == '(' {
		lexer.startToken(TokenTypeOpenBracket)
		lexer.endToken(char)
	} else if char == ')' {
		lexer.startToken(TokenTypeCloseBracket)
		lexer.endToken(char)
	} else if char == '"' {
		lexer.startToken(TokenTypeString)
		lexer.state = StateQuotedString
	} else {
		lexer.startToken(TokenTypeString)
		lexer.state = StateSimpleString
		lexer.buffer = char
	}
}

func (lexer *Lexer) handleSimpleStringState(char rune) {
	// TODO: look up what PHP says this checks.
	// if (ctype_space($char)) {
	// if char == ' ' {
	if unicode.IsSpace(char) {
		lexer.finalizeSimpleString()
	} else if char == '(' {
		lexer.finalizeSimpleString()
		lexer.startToken(TokenTypeOpenBracket)
		lexer.endToken(char)
	} else if char == ')' {
		lexer.finalizeSimpleString()
		lexer.startToken(TokenTypeCloseBracket)
		lexer.endToken(char)
	} else {
		lexer.buffer += char
	}
}

func (lexer *Lexer) handleQuotedStringState(char rune) {
	if char == '\\' {
		lexer.state = StateQuotedStringEscape
	} else if char == '"' {
		lexer.endToken(lexer.buffer)
		lexer.state = StateBegin
		lexer.buffer = ""
	} else {
		lexer.buffer += char
	}
}

func (lexer *Lexer) handleQuotedStringEscapeState(char rune) {
	lexer.state = StateQuotedString
	lexer.buffer += char
}

func (lexer *Lexer) finalizeSimpleString() {
	if strings.Compare("and", lexer.buffer) == 0 {
		lexer.nextToken.tokenType = TokenTypeLogicalAnd
	} else if strings.Compare("or", lexer.buffer) == 0 {
		lexer.nextToken.tokenType = TokenTypeLogicalOr
	} else if strings.Compare("not", lexer.buffer) == 0 {
		lexer.nextToken.tokenType = TokenTypeLogicalNot
	}

	lexer.endToken(lexer.buffer, -1)
	lexer.state = StateBegin
	lexer.buffer = ""
}

func (lexer *Lexer) startToken(tokeType TokenType) {
	lexer.nextToken = NewToken(
		tokenType,
		"",
		lexer.currentOffset,
		0,
		lexer.currentLine,
		lexer.currentColumn,
	)
}

func (lexer *Lexer) endToken(value string, lengthAdjustment int) {
	lexer.nextToken.Value = value
	lexer.nextToken.EndOffset = lexer.currentOffset + lengthAdjustment + 1
	lexer.tokens = append(lexer.tokens, lexer.nextToken)
	lexer.nextToken = nil
}
