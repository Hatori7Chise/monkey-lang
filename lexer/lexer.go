package lexer

import (
	"monkey-lang/token"
)

/*
	Lexer란
	소스코드를 입력받고 소스코드를 표현하는 토큰 열을 반환하는 행위를 맡은 놈을 뜻한다.
*/

/*
Lexer 구조체에서 입력 문자열을 가리키는 포인터(index)가 두개인 이유는
다음 처리 대상을 알아내려면 입력 문자열에서 다음 문자를 미리 살펴봄과 동시에 현재 문자를 보존하기 위함이다.
때문에 readPosition은 언제나 입력 문자열에서 다음 문자를 가리킨다.
또한 position은 입력 문자열에서 현재 문자를 가리킨다.

그리고 현재 문자가 곧 ch이다. 여기서 ch의 타입은 byte이다. 이 말인 즉슨 ASCLL만 지원한다는 뜻이며 UTF-8을 지원하기 위해선 rune 타입을 이용해야 한다.
만약 한글언어를 만든다고 한다면 rune 타입을 이용해야 한다는 것이다.
또한 UTF-8을 이용한다면 저장 방식이 byte와 다르기 때문에 밑에 정의된 함수들의 절차를 변경해야 한다. 저자는 이 부분을 연습문제로 남긴다고 한다.

또한 Lexer의 요소들은 Export를 하지 않는다.
*/
type Lexer struct {
	input        string // 입력
	position     int    // 입력에서 현재 위치 (현재 문자를 가리킨다.)
	readPosition int    // 입력에서 현재 읽는 위치 (현재 문자의 다음을 가리킨다.)	Lexer.input[lexcer.readPostition]
	ch           byte   // 현재 조사하고 있는 문자
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // position, readPosition, ch 초기화
	return l
}

/*
readChar는 export가 아니다.

readChar() 메서든는 다음 문자에 접근할 수 있게 만들어 준다. 즉, 문자열 input에서 렉서가 현재 보고 있는 위치에서
다음으로 이동하기 위한 메서드이다.

여기서 만약 문자열 끝에 도달하였다면, ASCLL 코드 문자 NULL에 해당하는 0을 넣는다. 그리고 0은 NULL의 뜻(아직 아무것도 읽지 않은 상태)도 있지만
EOF의 뜻(파일의 끝)도 있다.

여기서 중요한 점은 0과 문자 '0' 이다. 이것을 유의하며 코드를 보자
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { /*문자열의 끝에 도달했다면*/
		l.ch = 0 // EOF, NULL 대입
	} else { // 문자열 끝에 도달하지 않았다면 == 아직 읽어야할 문자열이 남았다면
		l.ch = l.input[l.readPosition] // 다음 문자열을 l.ch에 할당
	}

	l.position = l.readPosition // 현재 문자 index를 할당
	l.readPosition += 1         // 다음을 가리키는 readPostition의 index를 1 더해줌

}

// 현재 검사하고 있는 문자(lexer.ch)를 보고 이에 대응하는 토큰을 반환하는 함수이다.(요놈이 주 함수이다.)
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // Monkey-lang에서는 여백을 필요로 하지 않기 때문에 여백을 SKIP 한다

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0: // NULL || EOF
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) { // 글자이면? (여기서 글자라는 것은 a ~ z, A ~ Z, _를 말한다. 즉, 식별자와 예약어가 여기에 해당된다.)
			tok.Literal = l.readIdentifier() // 글자 읽기(이제 글자를 읽었으니, 이게 예약여 인지 식별자 인지 구분해야 한다.)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/*
isLetter 함수는 Export하지 않는다.

이 함수는 매우 직관적이면서 중요하다.
여기서 중요한 점은 _가 글자라고 인식하고 있다는 점이다.
만약 우리가 #를 글자라고 인식해야 한다면 함수 내부 조건식에 추가해줘야 한다.
*/
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

/*
여백을 skip하는 함수
WHY) 왜? 전체 코드를 가공하지 않고 skip하나?
ANW) 인터프리터 이기 때문에, 컴파이러 같은 경우 가공을 시켜 하는 방법도 좋을 텐데 지금 만들고 있는 것은 인터프리터 이기 때문에 skip을 해야 한다.
*/
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

/*
여기서 중요한 점은 진수 판별을 하지 않았다는 점이다. 만약 2진수, 8진수, 16진수를 지원하는 언어를 제작한다면
이에 대응되는 알고리즘을 작성해야만 한다.
*/
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}
