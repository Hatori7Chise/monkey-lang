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

/*
이젠 한 문자 토큰(-), 두 문자 토큰(==), 예약어 토큰(return)만 구현하면 된다. 여기서 중요한 점은
두 문자 토큰만 어떻게 구현할지 고민하면 된다.
*/
func TestNextTokenThreeCase(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
  		x + y;
	};

	let result = add(five, ten);


	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	`
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
		{token.LET, "let"}, // 10
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"}, // 20
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("}, // 30
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"}, // 40
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"}, // 50
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"}, // 60
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="}, // 66
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for index, element := range res {
		tok := l.NextToken()
		if tok.Type != element.expectedType {
			t.Fatalf("res[%d] - tokenType wrong. expeted=%q, got=%q", index, element.expectedType, tok.Type)
		}

		if tok.Literal != element.expectedLiteral {
			t.Fatalf("res[%d] = tokenLiteral wrong. expected=%q, got=%q", index, element.expectedLiteral, tok.Literal)
		}
	}

}
