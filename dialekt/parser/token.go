package dialekt

const TokenWildcardCharacter = "*"

type TokenType int

const (
	_ TokenType = iota
	TokenTypeLogicalAnd
	TokenTypeLogicalOr
	TokenTypeLogicalNot
	TokenTypeString
	TokenTypeOpenBracket
	TokenTypeCloseBracket
)

type Token struct {
	TokenType    TokenType
	Value        string
	StartOffset  int
	EndOffset    int
	LineNumber   int
	ColumnNumber int
}

func NewToken(tokenType TokenType, value string, startOffset, endOffset, lineNumber, columnNumber int) *Token {
	return &Token{
		tokenType,
		value,
		startOffset,
		endOffset,
		lineNumber,
		columnNumber,
	}
}

// TODO: Remove this, Pretty sure the String() code in next block below is what i want.
// func (token *Token) String() string {
// 	switch token.TokenType {
// 	case TokenTypeLogicalAnd:
// 		return "AND operator"
// 	case TokenTypeLogicalOr:
// 		return "OR operator"
// 	case TokenTypeLogicalNot:
// 		return "NOT operator"
// 	case TokenTypeString:
// 		return "tag"
// 	case TokenTypeOpenBracket:
// 		return "open bracket"
// 	case TokenTypeCloseBracket:
// 		return "close bracket"
// 	}

// 	panic("Unknown token type.")
// }

func (tokenType TokenType) String() string {
	switch token.TokenType {
	case TokenTypeLogicalAnd:
		return "AND operator"
	case TokenTypeLogicalOr:
		return "OR operator"
	case TokenTypeLogicalNot:
		return "NOT operator"
	case TokenTypeString:
		return "tag"
	case TokenTypeOpenBracket:
		return "open bracket"
	case TokenTypeCloseBracket:
		return "close bracket"
	}

	panic("Unknown token type.")
}
