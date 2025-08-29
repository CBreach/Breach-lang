package lexer

import (
	"testing"
	"slices"
	"github.com/Breach-lang/internal/token"
)

/*
	this tests if the lexer works when the rune encounter is a letter
	it should capture entire tokens:
		if var is passed, the lexer should handle var as a single token and not 3 separate ones
*/
func TestLexerdLetter(t *testing.T){
	expected := []token.Token{
		{Kind: token.KW_VAR, Lexeme: "var", LineNum: 1, ColNum: 1},
		{Kind: token.IDENT, Lexeme: "x", LineNum: 1, ColNum: 2},
	}
	token := lexer("var x")
	if !slices.Equal(token, expected){
		t.Fatal("Lexer mismatch")
	}
}

func TestLexerNumber(t *testing.T){
	expected := []token.Token{
		{Kind: token.NUMBER, Lexeme: "22", LineNum: 2, ColNum: 1}, //we make line number 2 because the test above updates the line count
		//although this is odd in isolated tests i belive this is actually the correct behavior since we'd technically be passing in consecutive lines 
	}
	token := lexer("22")
	if !slices.Equal(token, expected){
		t.Fatalf("Lexer mismatch: expected -> %v but got -> %v", expected, token)
	}
}
