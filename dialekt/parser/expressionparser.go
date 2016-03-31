package dialekt

import (
	"regex"
	"strings"
)

type ExpressionParser struct {
	AbstractParser
	logicalOrByDefault bool
}

func NewExpressionParser() *ExpressionParser {
	parser := &ExpressionParser{
		newAbstractParser(),
		false,
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

func (parser *ExpressionParser) parseExpression() (*AbstractExpression, error) {
	parser.startExpression()

	expression, err := parser.parseUnaryExpression()
	if err != nil {
		return nil, err
	}

	expression, err := parser.parseCompoundExpression(expression)
	if err != nil {
		return nil, err
	}

	parser.endExpression(expression)

	return expression
}

func (parser *ExpressionParser) parseUnaryExpression() (*AbstractParser, error) {
	foundExpected, err := parser.expectToken(
		TokenTypeString,
		TokenTypeLogicalNot,
		TokenTypeOpenBracket,
	)
	if !foundExpected {
		// I have the error available here (if one), might need to update return type.
		// The PHP version is using exceptions for control flow for this...
		return nil, err
	}

	if TokenTypeLogicalNot == parser.currentToken.TokenType {
		return parser.parseLogicalNot()
	} else if TokenTypeOpenBracket == parser.currentToken.TokenType {
		return parser.parseNestedExpression()
	} else if !strings.Contains(parser.wildcardString(), parser.currentToken.Value) {
		return parser.parseTag(), nil
	} else {
		return parser.parsePattern(), nil
	}
}

func (parser *ExpressionParser) parseTag() *AbstractExpression {
	parser.startExpression()

	expression := NewTag(
		parser.currentToken.Value,
	)

	parser.nextToken()

	parser.endExpression(expression)

	return expression
}

func (parser *ExpressionParser) parsePattern() (*AbstractExpression, error) {
	parser.startExpression()

	pattern := "/(" + regex.QuoteMeta(parser.wildcardString()) + ")/"
	re := regex.MustCompile(pattern)
	parts := re.Split(parser.currentToken.Value, -1)

	expression := NewPattern()

	for _, value := range parts {
		if parser.wildcardString() == value {
			expression.add(NewPatternWildcard())
		} else {
			expression.add(NewPatternLiteral(value))
		}
	}

	parser.nextToken()

	parser.endExpression(expression)

	return expressio, nil
}

func (parser *ExpressionParser) parseNestedExpression() (*AbstractExpression, error) {
	parser.startExpression()

	parser.nextToken()

	expression, err := parser.parseExpression()
	if err != nil {
		return nil, err
	}

	foundExpected, err := parser.expectToken(TokenTypeCloseBracket)
	if !foundExpected {
		// I have the error available here (if one), might need to update return type.
		// The PHP version is using exceptions for control flow for this...
		return nil, err
	}

	parser.nextToken()

	parser.endExpression(expression)

	return expressio, nil
}

func (parser *ExpressionParser) parseLogicalNot() (*AbstractExpression, error) {
	parser.startExpression()

	parser.nextToken()

	expression, err := parser.parseUnaryExpression()
	if err != nil {
		return nil, err
	}

	expression := NewLogicalNot(expression)

	parser.endExpression(expression)

	return expression, nil
}

func (parser *ExpressionParser) parseCompoundExpression(expression *ExpressionInterface, minimumPrecedence int) (*AbstractExpression, nil) {
	allowCollapse := false
	leftExpression := expression

	for {
		// Parse the operator and determine whether or not it's explicit ...
		operator, isExplicit = parser.parseOperator()

		precedence := operatorPrecedence(operator)

		// Abort if the operator's precedence is less than what we're looking for ...
		if precedence < minimumPrecedence {
			break
		}

		// Advance the token pointer if an explicit operator token was found ...
		if isExplicit {
			parser.nextToken()
		}

		// Parse the expression to the right of the operator ...
		rightExpression, err := parser.parseUnaryExpression()
		if err != nil {
			return nil, err
		}

		// Only parse additional compound expressions if their precedence is greater than the
		// expression already being parsed ...
		nextOperator, _ := parser.parseOperator()

		if precedence < operatorPrecedence(nextOperator) {
			rightExpression, er = parser.parseCompoundExpression(
				rightExpression,
				precedence+1,
			)

			if err != nil {
				return nil, err
			}
		}

		// Combine the parsed expression with the existing expression ...
		// Collapse the expression into an existing expression of the same type ...
		if oper == TokenTypeLogicalAnd {
			leftExpression, ok := leftExpression.(LogicalAnd)
			if allowCollapse && ok {
				leftExpression.add(rightExpression)
			} else {
				leftExpression = NewLogicalAnd(leftExpression, rightExpression)
				allowCollapse = true
			}
		} else if oper == TokenTypeLogicalOr {
			leftExpression, ok := leftExpression.(LogicalOr)
			if allowCollapse && ok {
				leftExpression.add(rightExpression)
			} else {
				leftExpression = NewLogicalOr(leftExpression, rightExpression)
				allowCollapse = true
			}
		} else {
			// return nil, error.New("Unknown operator type.")
			panic("Unknown operator type.")
		}
	}

	return leftExpression, nil
}

func (parser *ExpressionParser) parseOperator() (oper TokenType, isExplicit bool) {
	if currentToken == nil {
		// End of input ...
		return nil, false
	} else if TokenTypeCloseBracket == currentToken.TokenType {
		// Closing bracket ...
		return nil, false
	} else if TokenTypeLogicalOr == currentToken.TokenType {
		// Explicit logical OR ...
		return TokenTypeLogicalOr, true
	} else if TokenTypeLogicalAnd == currentToken.TokenType {
		// Explicit logical AND ...
		return TokenTypeLogicalAnd, true
	} else if logicalOrByDefault {
		// Implicit logical OR ...
		return TokenTypeLogicalOr, false
	} else {
		// Implicit logical AND ...
		return TokenTypeLogicalAnd, false
	}
}

func operatorPrecedence(oper TokenType) int {
	if oper == TokenTypeLogicalAnd {
		return 1
	} else if oper == TokenTypeLogicalOr {
		return 0
	} else {
		return -1
	}
}
