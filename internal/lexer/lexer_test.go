package lexer

import (
	"slices"
	"testing"

	"github.com/Breach-lang/internal/token"
)

/*
this tests if the lexer works when the rune encounter is a letter
it should capture entire tokens:

	if var is passed, the lexer should handle var as a single token and not 3 separate ones
*/
func TestLexerdLetter(t *testing.T) {
	expected := []token.Token{
		{Kind: token.KW_VAR, Lexeme: "var", LineNum: 1, ColNum: 1},
		{Kind: token.IDENT, Lexeme: "x", LineNum: 1, ColNum: 2},
	}
	token := lexer("var x")
	if !slices.Equal(token, expected) {
		t.Fatal("Lexer mismatch")
	}
}

func TestLexerNumber(t *testing.T) {
	expected := []token.Token{
		{Kind: token.NUMBER, Lexeme: "22", LineNum: 2, ColNum: 1}, //we make line number 2 because the test above updates the line count
		//although this is odd in isolated tests i belive this is actually the correct behavior since we'd technically be passing in consecutive lines
	}
	token := lexer("22")
	if !slices.Equal(token, expected) {
		t.Fatalf("Lexer mismatch: expected -> %v but got -> %v", expected, token)
	}
}
func TestLexer(t *testing.T) {
	expected := []token.Token{
		{Kind: token.KW_VAR, Lexeme: "var", LineNum: 3, ColNum: 1},
		{Kind: token.IDENT, Lexeme: "x", LineNum: 3, ColNum: 2},
		{Kind: token.ASSIGN, Lexeme: "=", LineNum: 3, ColNum: 3},
		{Kind: token.NUMBER, Lexeme: "24", LineNum: 3, ColNum: 4},
		{Kind: token.SCOLON, Lexeme: ";", LineNum: 3, ColNum: 5},
	}
	expected2 := []token.Token{
		{Kind: token.KW_IF, Lexeme: "if", LineNum: 4, ColNum: 1},
		{Kind: token.LPAREN, Lexeme: "(", LineNum: 4, ColNum: 2},
		{Kind: token.IDENT, Lexeme: "i", LineNum: 4, ColNum: 3},
		{Kind: token.LTE, Lexeme: "<=", LineNum: 4, ColNum: 4},
		{Kind: token.NUMBER, Lexeme: "2", LineNum: 4, ColNum: 5},
		{Kind: token.RPAREN, Lexeme: ")", LineNum: 4, ColNum: 6},
		{Kind: token.LBRACE, Lexeme: "{", LineNum: 4, ColNum: 7},
	}
	token := lexer("var x = 24;")
	token2 := lexer("if (i <= 2){")

	if !slices.Equal(token, expected) {
		t.Fatalf("Lexer mismatch: expected -> %v \n \t\t\t\t but got -> %v", expected, token)

	}
	if !slices.Equal(token2, expected2) {
		t.Fatalf("Lexer mismatch: expected -> %v \n \t\t\t\t but got -> %v", expected2, token2)

	}
}
func TestLexerStr(t *testing.T) {
	expected := []token.Token{
		{Kind: token.STRING, Lexeme: "this is a string", LineNum: 5, ColNum: 1},
	}
	token := lexer("\"this is a string\"")
	if !slices.Equal(token, expected) {
		t.Fatalf("Lexer mismatch: expected -> %v \n \t\t\t\t but got -> %v", expected, token)

	}
}
func TestLexerIllegalStr(t *testing.T) {
	expected := []token.Token{
		{Kind: token.ILLEGAL, Lexeme: "this is an invalid string", LineNum: 6, ColNum: 1},
	}
	token := lexer("\"this is an invalid string")

	if !slices.Equal(token, expected) {
		t.Fatalf("Lexer mismatch: expected -> %v \n \t\t\t\t but got -> %v", expected, token)

	}
}
func TestLexerIllegalToken(t *testing.T) {
	expected := []token.Token{
		{Kind: token.ILLEGAL, Lexeme: "_gato", LineNum: 7, ColNum: 1},
		{Kind: token.ASSIGN, Lexeme: "=", LineNum: 7, ColNum: 2},
		{Kind: token.NUMBER, Lexeme: "10", LineNum: 7, ColNum: 3},
	}
	token := lexer("_gato= 10")
	if !slices.Equal(expected, token) {
		t.Fatalf("Lexer mismatch: expected -> %v \n \t\t\t\t but got -> %v", expected, token)
	}
}
