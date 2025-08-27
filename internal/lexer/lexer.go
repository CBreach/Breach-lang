package lexer

import (
	"github.com/Breach-lang/internal/token"
)

func lexer(input string) []token.Token {
	input += "\n" // add a newline at the end of the input to make sure the last token is processed
	var tokens []token.Token

	current := 0      //counter for the current character in the input
	line, col := 1, 1 //track line and column numbers for error reporting
	/*
		this to me is still a bit confusing but basically, the weird syntax
		[]rune(input) converts the input string to a slice of runes (which are like characters but can represent more than just ASCII)
	*/
	runes := []rune(input)
	for current < len(runes) {
		currChar := runes[current] //stores the current character

		switch {
		case isLetter(currChar):
			// in the case that we encounter a letter it's important that we capture the entire thing as a single token
			var val string
			startPointer, endPointer := current, current

			for isAlphaNumeric(currChar) {
				col = current
				current++
				if current >= len(runes) {
					break
				}
				endPointer++
				currChar = runes[current]
			}
			val = string(runes[startPointer:endPointer])
			if key, ok := token.Keywords[val]; ok {
				tokens = append(tokens, buildToken(key, val, line, col))
			} else {
				tokens = append(tokens, buildToken(token.IDENT, val, line, col))
			}
		case isNumber(currChar):
			//TODO: handle floats
		default:
			// assign the token as being illegal

		}

	}
	return tokens
}
func buildToken(kind token.Kind, lexeme string, line int, col int) token.Token {
	return token.Token{
		Kind:    kind,
		Lexeme:  lexeme,
		LineNum: line,
		ColNum:  col,
	}
}
func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
func isNumber(ch rune) bool {
	return '0' <= ch && ch >= '0'
}
func isAlphaNumeric(ch rune) bool {
	return isLetter(ch) || isNumber(ch)
}
