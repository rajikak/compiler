package lexer

import (
	"testing"
	"token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
		      let ten = 10;
			  let add = fn(x,y){
				x+y;
			  };
  			  let result = add(five,ten);`

	tests := []struct {
		expectedType    token.Type
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
		{token.LPARN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPARN, ")"},
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
		{token.LPARN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPARN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literan wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
