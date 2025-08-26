package token

// defines the enum "go doesn't have built in enums so we have to work around"
type Kind string

const (
	LPAREN   Kind = "("
	RPAREN   Kind = ")"
	LBRACKET Kind = "["
	RBRACKET Kind = "]"
	LBRACE   Kind = "{"
	RBRACE   Kind = "}"
	PLUS     Kind = "+"
	MINUS    Kind = "-"
	STAR     Kind = "*"
	MOD      Kind = "%"
	SLASH    Kind = "/"
	BANG     Kind = "!"
	//two cahar symbols
	ASSIGN Kind = "="
	EQ     Kind = "=="
	NEQ    Kind = "!="
	GT     Kind = ">"
	LF     Kind = "<"
	GTE    Kind = ">="
	LTE    Kind = "<="

	//identifier
	INDENT Kind = "INDENT"
	NUMBER Kind = "NUMBER"
	STRING Kind = "STRING"

	//keywords
	KW_FUNC   Kind = "func"
	KW_VAR    Kind = "var"
	KW_LET    Kind = "let"
	KW_IF     Kind = "if"
	KW_ELSE   Kind = "else"
	KW_WHILE  Kind = "while"
	KW_FOR    Kind = "for"
	KW_DO     Kind = "Do"
	KW_IN     Kind = "in"
	KW_RETURN Kind = "return"
	KW_TRUE   Kind = "True"
	KW_FALSE  Kind = "False"
	KW_NIL    Kind = "nil"
	KW_OR     Kind = "or"
	KW_AND    Kind = "and"
	KW_IMPORT Kind = "import"

	EOF     Kind = "EOF"
	ILLEGAL Kind = "ILLEGAL"
)
