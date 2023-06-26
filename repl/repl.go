package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-lang/lexer"
	"monkey-lang/token"
)

/*
repl이란
read eval print loop 의 약자로
입력을 읽고(read), 인터프리터가 평가(eval)하고, 인터프리터가 결과물을 출력(print)하고, 이런 동작을 반복(loop)하는 놈을 뜻한다.
*/

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if scanned == false {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
