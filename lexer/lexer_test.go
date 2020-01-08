package lexer

import (
	"testing"
	"token"
)

func TestNextToken(t *testing.T) {
	inpit := `=+(){},;`

	tests :=[]struct {
		expectedType token.TokenType
		exepctedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""}
	}

	l := New(input)
	for i, test:= range tests {
		tok := l.NextToken()
		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, test.expectedType, test.Type)
		}

		if tok.Literal != test.exepctedLiteral {
			t.Fatalf("tests[%d] - literan wrong. expected=%q, got=%q", i, test.exepctedLiteral, test.Literal)
		}
	}
}
