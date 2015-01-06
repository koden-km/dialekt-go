package dialekt

const TokenWildcardCharacter = "*"

type TokenType int

const (
	TokenTypeLogicalAnd   TokenType = 1
	TokenTypeLogicalOr    TokenType = 2
	TokenTypeLogicalNot   TokenType = 3
	TokenTypeString       TokenType = 4
	TokenTypeOpenBracket  TokenType = 6
	TokenTypeCloseBracket TokenType = 7
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
	return &Token{tokenType, value, startOffset, endOffset, lineNumber, columnNumber}
}

func (token *Token) String() string {
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

	panic("Unknown type.")
}
