package test

import (
	"monkey-lang/lexer"
	"monkey-lang/token"
	"testing"
)

func TestNextTokenOneCase(t *testing.T) {
	input := `=+(){},;`

	res := []struct {
		expectedType    token.TokenType // 기대 TokenType
		expectedLiteral string          // 기대 Literal
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for index, element := range res {
		tok := l.NextToken()

		if tok.Type != element.expectedType {
			t.Fatalf("res[%d] - tokenType wrong. expected=%q, got=%q", index, element.expectedType, tok.Type)
		}

		if tok.Literal != element.expectedLiteral {
			t.Fatalf("res[%d] - literal wrong. expeted=%q, got=%q", index, element.expectedLiteral, tok.Literal)
		}
	}
}
