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

func TestNextTokenTwoCase(t *testing.T) {
	input := // "let five = 5; \n\tlet len = 10;\n\t\n\tlet add = fn(x, y) {\n\t\tx + y;\n\t};\n\t\n\tlet result = add(five, ten);\n\t"

		`let five = 5; 
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	`

	// 위 테스트 케이스 보다, 식별자, 에약어, 숫자가 존재한다.
	res := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for index, element := range res {
		var tok token.Token = l.NextToken()

		if tok.Type != element.expectedType {
			t.Fatalf("res[%d] - tokenType wrong. expected=%q, got=%q", index, element.expectedType, tok.Type)
		}

		if tok.Literal != element.expectedLiteral {
			t.Fatalf("res[%d] - tokenType wrong. expected=%q, got=%q", index, element.expectedLiteral, tok.Literal)
		}

	}

}
