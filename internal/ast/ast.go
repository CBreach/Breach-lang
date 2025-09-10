package ast

// lets begin by defining the base interfaces 


// Node is hte base of all AST nodes and provides a soruce span

type Node interface{
	Start() int
	End() int
}

type Expr interface{
	Node
	isExpr()
}
/*
	the methods isStmt() and isExpr() are simply tags to prevent cross-writting
	in other words they are not meant to be called, but serve as a way to distinguish
	categories that would otherwise look the same
*/
type Stmt interface{
	Node
	isStmt()
}


// instead of passing in the Start and End that are required by Node every time we can create a helper struct
type span struct{
	start, 
	end int
}
func (s span) Start() int{
	return  s.Start()
}
func (s span) End() int{
	return s.End()
}

/*Expressions*/

//Ident represents an identifier (variable/func name)
type Ident struct{
	span
	Name string
}

func (*Ident) isExpr(){} // so like mentioned this is just a marker method, in other words..
// its only purpose is to singal go that Ident is implementing the Expr interface

//LiteralKind distinguishes literal categories
type LiteralKind int
const(
	LitNumber LiteralKind = iota // so i just learned this.. iota automatically assigns a sequential incrementing integer to each constant (beginning at 0)
	LitString
	LitBool
	LitNil
)



// Literal holds a literal token
type Literal struct{
	span
	Kind LiteralKind
	Value string // this represents the raw lexme from the token
}
func (*Literal) isExpr(){} //again this does nothing but to tell the compiler that literal is implementing the Expr interface

//UnarayExpr represents prefix operators like !x or -x
type UnaryExpr struct{
	span
	Operator string
	exp Expr 
}
func (*UnaryExpr) isExpr(){}

// binary represents indix operators such as a+b, a==b and so on

type BinaryExpr struct{
	span
	Operator string
	Left, Right Expr
}
func (*BinaryExpr) isExpr(){}

//CallExpr represents function calls such as foo(a,b)
type CallExpr struct{
	span
	Callee Expr
	Args []Expr
}
func (*CallExpr) isExpr(){}


/*==========Statements!!==========*/

