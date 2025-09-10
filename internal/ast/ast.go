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

type Stmt interface{
	Node
	isStmt()
}


