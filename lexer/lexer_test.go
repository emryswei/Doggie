package lexer

import (
	"testing"
	"doggie/token"
)

func TestNextToken(t *testing.T){
	input := `=+(){},;`
	tests := []struct{
		expectedType     token.TokenType
		expectedLiteral  string
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

	l := New(input)

	for index, tt := range tests {
		tok := NextToken()
		if(tt.Type != tt.expectedType){
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",index, tt.expectedType, tok.Type)
		}
		if(tt.Literal != tt.expectedLiteral){
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",index,tt.expectedLiteral, tok.Literal)
		}
	}
	
}