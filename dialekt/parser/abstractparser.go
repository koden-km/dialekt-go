package dialekt

import (
	"strings"
)

type AbstractParser struct {
    wildcardString string
	tokenStack []*Token
	tokens []*Token
	tokenIndex int
	previousToken *Token

	currentToken *Token
}

func newAbstractParser() *AbstractParser {
	tokenStack := make([]*Token)
	tokens := make([]*Token)

	return &AbstractParser{TokenWildcardCharacter, tokenStack, tokens, 0, nil, nil}
}

// Fetch the string to use as a wildcard placeholder.
// Returns the string to use as a wildcard placeholder.
func (parser *AbstractParser) WildcardString() string {
	return parser.wildcardString
}

// Set the string to use as a wildcard placeholder.
// The string to use as a wildcard placeholder.
func (parser *AbstractParser) SetWildcardString(wildcardString string) string {
	parser.wildcardString = wildcardString
}

// The expression to parse.
// Returns the parsed expression or an error if the expression is invalid.
func (parser *AbstractParser) Parse(expression string, lexer *LexerInterface) (*ExpressionInterface, error) {
	if !lexer {
		lexer = NewLexer()
	}

	return parser.ParseTokens(
		lexer.lex(expression)
	)
}

// Parse an expression that has already beed tokenized.
// Returns the parsed expression or an error if the expression is invalid.
func (parser *AbstractParser) ParseTokens(tokens []*Token) (*ExpressionInterface, error) {
	if len(tokens) == 0 {
		return NewEmptyExpression(), nil
	}

	// TODO: is this the best way to clear an array/slice?
	parser.tokenStack = make([]*Token)
	parser.tokens = tokens
	parser.tokenIndex = 0
	parser.previousToken = nil
	parser.currentToken = tokens[0]

	var expression = parser.parseExpression()

	if parser.currentToken {
		// TODO: make a ParseError type?
		return nil, fmt.Errorf("Unexpected %s, expected end of input.", parser.currentToken.TokenType)
	}

	return expression, nil
}

// TODO: This might need to be public/exported?
func (parser *AbstractParser) parseExpression() (*AbstractExpression, error) {
	// TODO: panic or return error?
	// return nil, error.New("This method must be overridden")
	panic("This method must be overridden")
}

func (parser *AbstractParser) expectToken(types ...TokenType) (bool, error) {
	if !parser.currentToken {
		// TODO: make a ParseError type?
		return false, fmt.Errorf("Unexpected %s, expected %s.", parser.currentToken.TokenType, parser.formatExpectedTokenNames(types))
	} else {
		for _, tokenType := range types {
			if tokenType == parser.currentToken.TokenType {
				return true, nil
			}
		}
	}

	// TODO: make a ParseError type?
	return false, fmt.Errorf("Unexpected %s, expected %s.", parser.currentToken.TokenType, parser.formatExpectedTokenNames(types))
}

func (parser *AbstractParser) formatExpectedTokenNames(types ...TokenType) string {
	formattedTypes := make([]string, len(types))
	for i, tokenType := range types {
		formattedTypes[i] = tokenType.String()
	}

	if len(formattedTypes) == 1 {
		return formattedTypes[0]
	}

	lastType := formattedTypes[len(formattedTypes) - 1]

	return strings.Join(formattedTypes[:len(formattedTypes) - 1], ", ") + " or " + lastType;
}

func (parser *AbstractParser) nextToken() {
	parser.previousToken = parser.currentToken

	parser.tokenIndex++
	if parser.tokenIndex >= len(parser.tokens) {
		parser.currentToken = nil
	} else {
		parser.currentToken = parser.tokens[parser.tokenIndex]
	}
}

// Record the start of an expression.
func (parser *AbstractParser) startExpression() {
	parser.tokenStack = append(parser.tokenStack, parser.currentToken)
}

// Record the end of an expression.
func (parser *AbstractParser) endExpression(expression ExpressionInterface) {
	length := len(parser.tokenStack)
	lastToken = parser.tokenStack[length - 1]
	parser.tokenStack = parser.tokenStack[:length - 1]

	expression.setTokens(
		lastToken,
		parser.previousToken
	)
}
