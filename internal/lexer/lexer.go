package lexer

import (
	"errors"
	"log"

	"github.com/Breach-lang/internal/token"
	//"golang.org/x/text/currency"
)

// having this global here is probably not the best way of doing this...
// perhaphs making line a class variable for token and having a helper to increment would be a better approach..
// this will do for now
var line = 1 //TODO: encapsulate this so that its not a global variable
func lexer(input string) []token.Token { //perhaps it would be a good idea to pass the line counter as a variable to the function so that way different files start all at 0 when compiling
	input += "\n" // add a newline at the end of the input to make sure the last token is processed
	var tokens []token.Token

	current := 0 //counter for the current character in the input
	col := 1     //track line and column numbers for error reporting
	/*
		this to me is still a bit confusing but basically, the weird syntax
		[]rune(input) converts the input string to a slice of runes (which are like characters but can represent more than just ASCII)
	*/
	runes := []rune(input)
	for current < len(runes) {
		currChar := runes[current] //stores the current character
		switch {
		case token.TokenTable[string(currChar)] != "":
			if currChar == '!' || currChar == '<' || currChar == '>' || currChar == '=' {
				tokens = append(tokens, checkNextToken(runes, &current, col))
			} else {
				tokVal := string(currChar)
				tokKind, ok := token.TokenTable[tokVal]
				if !ok { //this is a bit redundant since this should technically never happen
					log.Fatalf("Could not find token %s", tokVal)
				}
				tokens = append(tokens, buildToken(tokKind, tokVal, line, col))
				current++
			}
		case isLetter(currChar, current):
			// in the case that we encounter a letter it's important that we capture the entire thing as a single token
			var val string
			startPointer, endPointer := current, current

			endPointer, err := traverseToken(runes, "letter", startPointer)
			current = endPointer + 1 // since we moved the token traversal to a independent function we have to update current manually
			val = string(runes[startPointer:endPointer])
			if err != nil { //if we have an error we know the token is invalid
				tokens = append(tokens, buildToken(token.ILLEGAL, val, line, col))
			}
			if key, ok := token.Keywords[val]; ok {
				tokens = append(tokens, buildToken(key, val, line, col))
			} else {
				tokens = append(tokens, buildToken(token.IDENT, val, line, col))
			}
		case isNumber(currChar):
			// similar to how we handled letters, when we encounter a number we want to grab the hole thing as a single token rather
			//TODO: handle floats
			var val string
			startPointer := current
			endPointer, err := traverseToken(runes, "number", startPointer)
			current = endPointer
			val = string(runes[startPointer:endPointer])
			if err != nil {
				tokens = append(tokens, buildToken(token.ILLEGAL, val, line, col))
			} else {
				tokens = append(tokens, buildToken(token.NUMBER, val, line, col))
			}

		case currChar == ' ':
			current++
			continue
		case currChar == '\n':
			current++
			continue
		case currChar == '/':
			if current+1 < len(runes) && runes[current+1] == '/' {
				// we have a comment, we want to skip everything until the end of the line
				for current < len(runes) && runes[current] != '\n' {
					current++
				}
				continue
			} else {
				tokens = append(tokens, buildToken(token.SLASH, string(currChar), line, col))
				current++
			}
		case currChar == '"':
			//similar to how we handle letters and numbers we want to grab the entire string as a single token
			startPointer := current + 1 //skip the opening quote
			endPointer, err := traverseToken(runes, "string", startPointer)
			if err != nil {
				tokens = append(tokens, buildToken(token.ILLEGAL, string(runes[startPointer:endPointer-1]), line, col))
				current = endPointer
			} else {
				current = endPointer + 1 //move past the closing quote
				tokens = append(tokens, buildToken(token.STRING, string(runes[startPointer:endPointer]), line, col))
			}
		default:
			// assign the token as being illegal
			// grab from first index of the token up until the char before the next space or \n character
			invalidTok := handleInvalidToken(runes, &current)
			tokens = append(tokens, buildToken(token.ILLEGAL, invalidTok, line, col))
		}
		col++
	}
	line++
	return tokens
}
func handleInvalidToken(runes []rune, current *int) string {
	currentChar := string(runes[*current])
	startptr := *current
	for *current < len(runes) && token.TokenTable[currentChar] == "" {
		*current++
		currentChar = string(runes[*current])
	}
	return string(runes[startptr:*current])

}
func checkNextToken(runes []rune, current *int, col int) token.Token {
	tokVal := string(runes[*current])
	if *current+1 < len(runes) && runes[*current+1] == '=' {
		tokVal = string(runes[*current : *current+2])
		*current += 2
	} else {
		*current++
	}
	tokKind, ok := token.TokenTable[tokVal]
	if !ok {
		log.Fatalf("Token not found: %s", tokKind)
	}
	return buildToken(tokKind, tokVal, line, col)
}
func traverseToken(runes []rune, tokenType string, strIndex int) (int, error) {
	endIndex := strIndex
	current := strIndex
	currChar := runes[current]
	if tokenType == "letter" {
		for isAlphaNumeric(currChar, current) || currChar == '\n' {
			current++
			if current >= len(runes) {
				break
			}
			endIndex++
			currChar = runes[current]
		}

	} else if tokenType == "number" { //when its a number we wan't to loop for the entire digit not excluding decimal points
		decimalCount := 0
		for isNumber(currChar) || currChar == '.' { //we can safelly do this because we know that at least the first char is a number
			if currChar == '.' {
				decimalCount++
			}
			current++
			if current >= len(runes) {
				break
			}
			endIndex++
			currChar = runes[current]
		}
		if decimalCount > 1 {
			return endIndex, errors.New("invalid token... multiple decimal points")
		}
	} else if tokenType == "string" {
		// we want to keep going until we find the closing quote
		for currChar != '"' && current < len(runes) {
			current++
			endIndex++
			if current >= len(runes) {
				return endIndex, errors.New("unterminated string")
			}
			currChar = runes[current]
		}
	}
	return endIndex, nil
}
func buildToken(kind token.Kind, lexeme string, line int, col int) token.Token {
	return token.Token{
		Kind:    kind,
		Lexeme:  lexeme,
		LineNum: line,
		ColNum:  col,
	}
}
func isLetter(ch rune, current int) bool {
	if current == 0 && ch == '_' {
		return false //we do this because no token can begin with _
	}
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9' || ch >= 48 && ch <= 57
}
func isAlphaNumeric(ch rune, current int) bool {
	return isLetter(ch, current) || isNumber(ch)
}
