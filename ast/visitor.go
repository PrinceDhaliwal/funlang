package ast

type Visitor interface {
	VisitNilLiteral(*NilLiteral)
	VisitNumericLiteral(*NumericLiteral)
	VisitStringLiteral(*StringLiteral)
	VisitBooleanLiteral(*BooleanLiteral)
	VisitIdentifier(*Identifier)
	VisitArgumentList(*ArgumentList)
	VisitMemberExpression(*MemberExpression)
	VisitPrefixExpression(*PrefixExpression)
	VisitPostfixExpression(expression *PostfixExpression)
	VisitBinaryExpression(*BinaryExpression)
	VisitAssignExpression(*AssignExpression)
	VisitArrayType(*ArrayType)
	VisitField(*Field)
	VisitStructType(*StructType)
	VisitFuncType(*FuncType)
	VisitDeclaration(*Declaration)
	VisitDeclarationList(*DeclarationList)
	VisitTypeDeclaration(*TypeDeclaration)
	VisitBlockStatement(*BlockStatement)
	VisitForStatement(*ForStatement)
	VisitExpressionStmt(*ExpressionStmt)
	VisitFunctionProtoType(*FunctionProtoType)
	VisitFunctionStatement(*FunctionStatement)
	VisitIfElseStatement(statement *IfElseStatement)
	VisitReturnStatement(statement *ReturnStatement)
}

type VisitorAcceptor interface {
	Accept(Visitor)
}
