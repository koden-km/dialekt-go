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

// TODO: Up to here

// private function handleBeginState($char)
// {
// 	if (ctype_space($char)) {
// 		// ignore ...
// 	} elseif ($char === '(') {
// 		$this->startToken(Token::OPEN_BRACKET);
// 		$this->endToken($char);
// 	} elseif ($char === ')') {
// 		$this->startToken(Token::CLOSE_BRACKET);
// 		$this->endToken($char);
// 	} elseif ($char === '"') {
// 		$this->startToken(Token::STRING);
// 		$this->state = self::STATE_QUOTED_STRING;
// 	} else {
// 		$this->startToken(Token::STRING);
// 		$this->state = self::STATE_SIMPLE_STRING;
// 		$this->buffer = $char;
// 	}
// }

// private function handleSimpleStringState($char)
// {
// 	if (ctype_space($char)) {
// 		$this->finalizeSimpleString();
// 	} elseif ($char === '(') {
// 		$this->finalizeSimpleString();
// 		$this->startToken(Token::OPEN_BRACKET);
// 		$this->endToken($char);
// 	} elseif ($char === ')') {
// 		$this->finalizeSimpleString();
// 		$this->startToken(Token::CLOSE_BRACKET);
// 		$this->endToken($char);
// 	} else {
// 		$this->buffer .= $char;
// 	}
// }

// private function handleQuotedStringState($char)
// {
// 	if ($char === '\\') {
// 		$this->state = self::STATE_QUOTED_STRING_ESCAPE;
// 	} elseif ($char === '"') {
// 		$this->endToken($this->buffer);
// 		$this->state = self::STATE_BEGIN;
// 		$this->buffer = '';
// 	} else {
// 		$this->buffer .= $char;
// 	}
// }

// private function handleQuotedStringEscapeState($char)
// {
// 	$this->state = self::STATE_QUOTED_STRING;
// 	$this->buffer .= $char;
// }

// private function finalizeSimpleString()
// {
// 	if (strcasecmp('and', $this->buffer) === 0) {
// 		$this->nextToken->type = Token::LOGICAL_AND;
// 	} elseif (strcasecmp('or', $this->buffer) === 0) {
// 		$this->nextToken->type = Token::LOGICAL_OR;
// 	} elseif (strcasecmp('not', $this->buffer) === 0) {
// 		$this->nextToken->type = Token::LOGICAL_NOT;
// 	}

// 	$this->endToken($this->buffer, -1);
// 	$this->state = self::STATE_BEGIN;
// 	$this->buffer = '';
// }

// private function startToken($type)
// {
// 	$this->nextToken = new Token(
// 		$type,
// 		'',
// 		$this->currentOffset,
// 		0,
// 		$this->currentLine,
// 		$this->currentColumn
// 	);
// }

// private function endToken($value, $lengthAdjustment = 0)
// {
// 	$this->nextToken->value = $value;
// 	$this->nextToken->endOffset = $this->currentOffset
// 								+ $lengthAdjustment
// 								+ 1;
// 	$this->tokens[] = $this->nextToken;
// 	$this->nextToken = null;
// }
