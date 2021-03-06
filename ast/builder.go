package ast

import (
	"fmt"

	"funlang/context"
	"funlang/lex"
)

type Builder struct {
	ctx *context.Context
}

func NewBuilder(ctx *context.Context) *Builder {
	return &Builder{ctx}
}

func (b *Builder) NewNilLiteral(pos lex.Position) *NilLiteral {
	return &NilLiteral{pos}
}

func (b *Builder) NewNumericLiteral(pos lex.Position, val string, isFloat bool) *NumericLiteral {
	return &NumericLiteral{pos: pos, val: val, isFloating: isFloat}
}

func (b *Builder) NewInt(num int) *NumericLiteral {
	return &NumericLiteral{pos: lex.NO_POS, val: fmt.Sprint(num), isFloating: false}
}

func (b *Builder) NewStringLiteral(pos lex.Position, val string) *StringLiteral {
	return &StringLiteral{pos: pos, val: val}
}

func (b *Builder) NewBooleanLiteral(pos lex.Position, val bool) *BooleanLiteral {
	return &BooleanLiteral{pos, val}
}

func (b *Builder) NewIdentifier(pos lex.Position, name string) *Identifier {
	return &Identifier{pos, name, nil}
}

func (b *Builder) NewArgumentList(pos lex.Position, args []Expression) *ArgumentList {
	return &ArgumentList{pos, args}
}

func (b *Builder) NewMemberExpression(token lex.TokenType, x Expression, member Expression) *MemberExpression {
	return &MemberExpression{token: token, member: member, x: x}
}

func (b *Builder) NewPrefixExpression(pos lex.Position, op lex.TokenType,
	x Expression) *PrefixExpression {
	return &PrefixExpression{pos: pos, op: op, x: x}
}

func (b *Builder) NewPostfixExpression(pos lex.Position, op lex.TokenType,
	x Expression) *PostfixExpression {
	return &PostfixExpression{pos: pos, op: op, x: x}
}

func (b *Builder) NewBinaryExpression(pos lex.Position, tok lex.TokenType, left, right Expression) *BinaryExpression {
	return &BinaryExpression{pos: pos, op: tok, left: left, right: right}
}

func (b *Builder) NewAssignExpression(pos lex.Position, left, right Expression) *AssignExpression {
	return &AssignExpression{pos: pos, left: left, right: right}
}

func (b *Builder) NewArrayType(pos lex.Position, size, t Expression) *ArrayType {
	return &ArrayType{pos: pos, size: size, t: t}
}

func (b *Builder) NewField(name, t Expression) *Field {
	return &Field{name: name, t: t}
}

func (b *Builder) NewStructType(pos lex.Position, fields []*Field) *StructType {
	return &StructType{pos: pos, fields: fields}
}

func (b *Builder) NewFuncType(pos lex.Position, params []Expression, ret Expression) *FuncType {
	return &FuncType{pos: pos, params: params, ret: ret}
}

func (b *Builder) NewDeclaration(pos lex.Position, name string, t Expression,
	init Expression) *Declaration {
	return &Declaration{pos: pos, name: name, t: t, init: init}
}

func (b *Builder) NewDeclarationList(pos lex.Position, decls []*Declaration) *DeclarationList {
	return &DeclarationList{pos: pos, decls: decls}
}

func (b *Builder) NewTypeDeclaration(pos lex.Position, name *Identifier, t Expression) *TypeDeclaration {
	return &TypeDeclaration{pos: pos, name: name, t: t}
}

func (b *Builder) NewFunctionProtoType(pos lex.Position, name string, args []DeclNode, ret Expression) *FunctionProtoType {
	return &FunctionProtoType{pos: pos, name: name, args: args, ret: ret}
}

func (b *Builder) NewFunctionStatement(proto *FunctionProtoType, body *BlockStatement) *FunctionStatement {
	return &FunctionStatement{proto: proto, body: body}
}

func (b *Builder) NewIfStatement(condition Expression, body Statement, els Statement) *IfElseStatement {
	return &IfElseStatement{condition: condition, body: body, elseNode: els}
}

func (b *Builder) NewForStatement(pos lex.Position, init, condition Expression, body Statement) *ForStatement {
	return &ForStatement{pos: pos, init: init, condition: condition, body: body}
}

func (b *Builder) NewExpressionStatement(expr Expression) *ExpressionStmt {
	return &ExpressionStmt{expr: expr}
}

func (b *Builder) NewReturnStatement(pos lex.Position, expr Expression) *ReturnStatement {
	return &ReturnStatement{pos: pos, expr: expr}
}

func (b *Builder) NewBlockStatement(list []Statement) *BlockStatement {
	return &BlockStatement{stmts: list}
}

func (b *Builder) NewDeclarationStatement(decl DeclNode) *DeclarationStatement {
	return &DeclarationStatement{decl}
}
