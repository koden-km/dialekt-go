package dialekt

import (
	"strings"
	"strconv"
	"regex"
)

type ExpressionParser struct {
	AbstractParser
	logicalOrByDefault bool
}

func NewExpressionParser() *ExpressionParser {
	parser := &ExpressionParser{
		newAbstractParser(),
		false
	}

	return parser
}

// Indicates whether or not the the default operator should be OR, rather than AND.
// Returns true if the default operator should be OR, rather than AND.
func (parser *ExpressionParser) LogicalOrByDefault() bool {
	return parser.logicalOrByDefault
}

// Set whether or not the the default operator should be OR, rather than AND.
// Set to true if the default operator should be OR, rather than AND.
func (parser *ExpressionParser) SetLogicalOrByDefault(logicalOrByDefault bool) {
	parser.logicalOrByDefault = logicalOrByDefault
}

func (parser *ExpressionParser) parseExpression() *AbstractExpression {
	parser.startExpression()

	expression := parser.parseUnaryExpression()
	expression = parser.parseCompoundExpression(expression)

	parser.endExpression(expression)

	return expression
}

func (parser *ExpressionParser) parseUnaryExpression() *AbstractParser {
	foundExpected, err := parser.expectToken(
		TokenTypeString,
		TokenTypeLogicalNot,
		TokenTypeOpenBracket
	)
	if !foundExpected {
		// I have the error available here (if one), might need to update return type.
		// The PHP version is using exceptions for control flow for this...
		return nil
	}

	if TokenTypeLogicalNot == parser.currentToken.type {
		return parser.parseLogicalNot()
	} else if TokenTypeOpenBracket == $this.currentToken.type {
		return parser.parseNestedExpression()
	} else if !strings.Contains(parser.wildcardString(), parser.currentToken.value) {
		return parser.parseTag()
	} else {
		return parser.parsePattern()
	}
}

func (parser *ExpressionParser) parseTag() *AbstractExpression {
	parser.startExpression()

	expression := NewTag(
		parser.currentToken.value
	)

	parser.nextToken()

	parser.endExpression(expression)

	return expression
}

func (parser *ExpressionParser) parsePattern() *AbstractExpression {
	parser.startExpression()


// HOW TO
	// check what preg_quote() 2nd arg is. Is it the only thing to quote?
	// '/(' . preg_quote(parser.wildcardString(), '/') . ')/',
	pattern := "/(" + strconv.Quote(parser.wildcardString()) + ")/"

	re := regex.MustCompile(pattern)

// TODO use re. one of these...

// func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int {
// func (re *Regexp) Split(s string, n int) []string {

// END


	parts := preg_split(
		'/(' . preg_quote(parser.wildcardString(), '/') . ')/',
		parser.currentToken->value,
		-1,
		PREG_SPLIT_DELIM_CAPTURE | PREG_SPLIT_NO_EMPTY
	)


	expression = NewPattern()

	for _, value := range parts {
		if parser.wildcardString() == value {
			expression.add(NewPatternWildcard())
		} else {
			expression.add(NewPatternLiteral(value))
		}
	}

	parser.nextToken()

	parser.endExpression(expression)

	return expression
}

