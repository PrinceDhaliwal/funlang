package ast

import (
	"fmt"
	"funlang/context"
	"funlang/lex"
	"strings"
)

type Program struct {
	source lex.Source
	ctx    *context.Context
	decls  []Node
}

func (p *Program) Beg() lex.Position { return lex.Position{0, 0} }
func (p *Program) End() lex.Position { return lex.Position{0, 0} }

func NewProgram(ctx *context.Context, source lex.Source, decls []Node) *Program {
	return &Program{source: source, ctx: ctx, decls: decls}
}

func (p *Program) Decls() []Node {
	return p.decls
}

func (p *Program) String() string {
	builder := strings.Builder{}
	for _, decl := range p.decls {
		builder.WriteString(fmt.Sprint(decl))
		builder.WriteString("\n")
	}

	return builder.String()
}
