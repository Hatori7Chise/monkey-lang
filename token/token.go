package token

/*
일번적인 코드는 enum을 이용해 미리 정의해 제한된 값을 사용하게 되는데
TokenType을 만들게 된다면, 개발자가 필요한 만큼 언제든지 TokenType을 추가할 수 있다.
int나 byte가 성능상 이점을 가지게 되지만, monkey-lang은 공부가 목적이므로 string 형으로 대체한다.
*/
type TokenType string

type Token struct {
	Type    TokenType // 렉서나 파서가 TokenType을 보고 판단한다.
	Literal string    // 값을 보존시킨다.
}

// TokenType 상수 정의
const (
	ILLEGAL = "ILLEGAL" // 한국어로 불법적인 뜻을 가짐 (monkey-lang에서는 어떤 토큰이나 문자가 렉서가 알 수 없다는 뜻으로 사용된다.)
	EOF     = "EOF"     // End of FILE(monkey-lang 에서는 파서에게 이제 그만 멈춰도 좋다라고 말하는 용도로 사용된다.)

	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"

	// 연산자
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
