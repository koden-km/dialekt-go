package dialekt

type AbstractParser struct {
    wildcardString string
	tokenStack []*Token
	tokens []*Token
	tokenIndex int
	previousToken *Token

	currentToken *Token

	// TODO: needs something like this:
	// abstract protected function parseExpression();
	parseExpressionFunc func() (*AbstractExpression, error)
}

func newAbstractParser() *Token {
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

// Parse an expression.
//
// @param string         $expression The expression to parse.
// @param LexerInterface $lexer      The lexer to use to tokenise the string, or nil to use the default.
//
// @return ExpressionInterface The parsed expression.
// @throws ParseException      if the expression is invalid.
func (parser *AbstractParser) Parse(expression string, lexer *LexerInterface) (*ExpressionInterface, error) {
	if !lexer {
		lexer = NewLexer()
	}

	return parser.ParseTokens(
		lexer.lex(expression)
	)
}

// Parse an expression that has already beed tokenized.
//
// @param array<Token> The array of tokens that form the expression.
//
// @return ExpressionInterface The parsed expression.
// @throws ParseException      if the expression is invalid.
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

	if !parser.parseExpressionFunc {
		// error...
		panic("Parser parse expression func not defined.")
	}

	var expression = parser.parseExpressionFunc()

	if parser.currentToken {
		// TODO: make a ParseError type?
		return nil, fmt.Errorf("Unexpected %s, expected end of input.", parser.currentToken.TokenType)
	}

	return expression, nil
}

func (parser *AbstractParser) expectToken(types ...TokenType) (bool, error) {
	if !parser.currentToken {
		// TODO: make a ParseError type?
		return false, fmt.Errorf("Unexpected %s, expected %s.", parser.currentToken.TokenType, parser.formatExpectedTokenNames(types))
// UP TO HERE
	} else {
		for _, tokenType := range types {
			if tokenType != parser.currentToken.TokenType {
				// TODO: how to check if something is in an array?
				return false, fmt.Errorf("Unexpected %s, expected %s.", parser.currentToken.TokenType, parser.formatExpectedTokenNames(types))
			}
		}
	}

	return true, nil



	// $types = func_get_args();

	// if (!parser.currentToken) {
	// 	throw new ParseException(
	// 		'Unexpected end of input, expected ' . parser.formatExpectedTokenNames($types) . '.'
	// 	);
	// } elseif (!in_array(parser.currentToken->type, $types)) {
	// 	throw new ParseException(
	// 		'Unexpected ' . Token::typeDescription(parser.currentToken->type) . ', expected ' . parser.formatExpectedTokenNames($types) . '.'
	// 	);
	// }
}

func (parser *AbstractParser) formatExpectedTokenNames(array $types) {
	// $types = array_map(
	// 	'Icecave\Dialekt\Parser\Token::typeDescription',
	// 	$types
	// );

	// if (count($types) === 1) {
	// 	return $types[0];
	// }

	// $lastType = array_pop($types);

	// return implode(', ', $types) . ' or ' . $lastType;
}

func (parser *AbstractParser) nextToken() {
	// parser.previousToken = parser.currentToken

	// if ++parser.tokenIndex >= len(parser.tokens) {
	// 	parser.currentToken = nil
	// } else {
	// 	parser.currentToken = parser.tokens[parser.tokenIndex]
	// }
}

/**
 * Record the start of an expression.
 */
func (parser *AbstractParser) startExpression() {
	// parser.tokenStack[] = parser.currentToken;
}

/**
 * Record the end of an expression.
 *
 * @return ExpressionInterface
 */
func (parser *AbstractParser) endExpression(ExpressionInterface $expression) {
	// $expression->setTokens(
	// 	array_pop(parser.tokenStack),
	// 	parser.previousToken
	// );
}
