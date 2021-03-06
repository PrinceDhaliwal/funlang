package ast

import "fmt"

type Visitor interface {
	Visit(node Node) Visitor
}

// this code is copied (kind of) from go/ast/walk.go
func Walk(v Visitor, node Node) {
	if v = v.Visit(node); v == nil {
		return
	}

	switch n := node.(type) {
	case *ArgumentList:
		for _, expr := range n.exprs {
			Walk(v, expr)
		}

	case *MemberExpression:
		Walk(v, n.member)
		Walk(v, n.x)

	case *PrefixExpression:
		Walk(v, n.x)

	case *PostfixExpression:
		Walk(v, n.x)

	case *BinaryExpression:
		Walk(v, n.left)
		Walk(v, n.right)

	case *AssignExpression:
		Walk(v, n.left)
		Walk(v, n.right)

	case *ArrayType:
		Walk(v, n.t)

	case *StructType:
		for _, field := range n.fields {
			Walk(v, field)
		}

	case *FuncType:
		for _, param := range n.params {
			Walk(v, param)
		}

		Walk(v, n.ret)

	case *Declaration:
		if n.t != nil {
			Walk(v, n.t)
		}
		if n.init != nil {
			Walk(v, n.init)
		}

	case *DeclarationList:
		for _, decl := range n.decls {
			Walk(v, decl)
		}

	case *TypeDeclaration:
		Walk(v, n.name)
		Walk(v, n.t)

	case *BlockStatement:
		for _, stmt := range n.stmts {
			Walk(v, stmt)
		}

	case *ForStatement:
		Walk(v, n.init)
		Walk(v, n.condition)
		Walk(v, n.body)

	case *FunctionProtoType:
		for _, arg := range n.args {
			Walk(v, arg)
		}

		Walk(v, n.ret)

	case *FunctionStatement:
		Walk(v, n.proto)
		Walk(v, n.body)

	case *ExpressionStmt:
		Walk(v, n.expr)

	case *IfElseStatement:
		Walk(v, n.condition)
		Walk(v, n.body)
		if n.elseNode != nil {
			Walk(v, n.elseNode)
		}

	case *ReturnStatement:
		Walk(v, n.expr)

	case *DeclarationStatement:
		Walk(v, n.decl)

	case *Field:
		Walk(v, n.name)
		Walk(v, n.t)

	case *Program:
		for _, decl := range n.decls {
			Walk(v, decl)
		}

	case *NumericLiteral:
	case *StringLiteral:
	case *BooleanLiteral:
		func(){}()
	default:
		panic(fmt.Sprintf("didn't expect this type, %T", n))
	}

}

type VisitorAcceptor interface {
	Accept(Visitor)
}
