### Q) Golang은 강타입 언어입니다. 하지만 CONST로 토큰들을 선언했였고, 이것을 TokenType에 할당했습니다. 이게 가능한건가요?
```go
res := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN,    "="},
		{token.PLUS,      "+"},
		{token.LPAREN,    "("},
		{token.RPAREN,    ")"},
		{token.LBRACE,    "{"},
		{token.RBRACE,    "}"},
		{token.COMMA,     ","},
		{token.SEMICOLON, ";"},
		{token.EOF,       "" },
	}
```

### A) Golnag에서는 상수를 Type와 Untype로 구분합니다.

여기서 Typed 상수는 선언에 형식을 지정한 상수를 뜻합니다. 
```go
const typeConst int32 = 8
```
위 상수는 int32 타입의 변수에만 값을 할당시킬 수 있습니다.

그러면 Untyped 상수는 무엇일까요?

위에서 아마 감을 잡았을 것 같은것 같습니다. 바로 타입이 지정되지 않은 상수를 뜻합니다.

예를 들면

```go
123        //Default hidden type is int
"circle"   //Default hidden type is string
5.6.       //Default hidden type is float64
true       //Default hidden type is bool
'a'        //Default hidden type is rune
3+5i       //Default hidden type is complex128
```

```go
const a = 123        //Default hidden type is int
const b = "circle"   //Default hidden type is string
const c = 5.6       //Default hidden type is float64
const d = true       //Default hidden type is bool
const e = 'a'        //Default hidden type is rune
const f = 3+5i       //Default hidden type is complex128
```

여기서 hidden type이란 입력되지 않은 상수에게 주어진 하나의 타입입니다. (아무튼 타입을 하나 지정해 줘야지 변수로서 존재할 수 있기 때문에 그렇것 같다. 추측입니다^^;)


상수의 기본 숨김 유형

|보통 우리가 나누는 타입|Golang Type|
|---|---|
|정수|int|
|실수|float64|
|복소수|complex128|
|문자열|string|
|불린형|bool|
|문자형|int32 or rune|

    참고: rune의 int32의 별칭이다. 

그리고 untyped 상수를 이용하면 type 갈래에 벗어나지 않으면 값을 대입할 수 있다. 

그러므로 untyped 상수는 hidden type을 가지고 보통 우리가 나누는 타입 갈래를 벗어나지 않은 이상 값을 할당 시킬 수 있다.

예를 들면

Math 패키지의 Pi의 값은 이렇게 정의되어 있다.
```go
const Pi = 3.14159265358979323846264338327950288419716939937510582097494459
```

그리고 사용은 이렇게 할 수 있다.

```go
package main
import (
    "fmt"
    "math"
)
func main() {
    var f1 float32
    var f2 float64
    f1 = math.Pi
    f2 = math.Pi

    fmt.Printf("Type: %T Value: %v\n", math.Pi, math.Pi)
    fmt.Printf("Type: %T Value: %v\n", f1, f1)
    fmt.Printf("Type: %T Value: %v\n", f2, f2)
}
```

그러므로 질문에서 참조한 코드는 아무 이상 돌아가는 것이 정상이다.

참고 사이트 : https://golangbyexample.com/typed-untyped-constant-golang/

### Q) fmt.Printf 함수의 인자로 넣을 수 있는 서식문자 중 %q는 무엇인가요?
```go
fmt.Printf("%q", "Hi\n")
```

### A) %q는 작은따옴표로 묶인 문자 리터럴을 Go 구문으로 안전하게 이스케이프합니다.
여기서 이스케이프 한다는 것은 그냥 출력한다는 뜻인것 같다. 
위 예제를 출력하면 그대로 출력된다. 
```
output: Hi\n
```