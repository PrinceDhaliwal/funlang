package ssa

import (
	"fmt"
	"funlang/types"
	"strings"
)

type Argument struct {
	valueWithName
	valueWithUsers
	t types.Type
}

func (a *Argument) Type() types.Type {
	return a.t
}

func (a *Argument) Tag() ValueTag {
	return ARGUMENT
}

func (a *Argument) Uses() []Value {
	return []Value{}
}

func (a *Argument) String() string {
	return fmt.Sprintf("%s:%s", a.Name(), a.Type())
}

func (a *Argument) ShortString() string {
	return a.String()
}

type Function struct {
	valueWithName
	valueWithUsers

	Blocks []*BasicBlock
	t      types.Type
	Args   map[string]*Argument

	// these fields are required while creating ssa
	current *BasicBlock
	locals  map[string]Value
	Extern  bool

	// for names
	nameCount int
}

func (f *Function) Uses() []Value {
	return []Value{}
}

func (f *Function) Tag() ValueTag {
	return FUNCTION
}

func (f *Function) Type() types.Type {
	return f.t
}

func (f *Function) ShortString() string {
	return fmt.Sprintf("%%%s", f.Name())
}

func (f *Function) GetArg(name string) *Argument {
	return f.Args[name]
}

func (f *Function) NextName() string {
	f.nameCount = f.nameCount + 1
	return fmt.Sprintf("t%d", f.nameCount)
}

func (f *Function) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%s %s(", types.ToFunctionType(f.t).ReturnType(), f.Name()))
	l := len(f.Args)
	i := 0
	for _, arg := range f.Args {
		builder.WriteString(arg.String())
		if i+1 == l {
			break
		}
		i++
		builder.WriteString(", ")
	}

	if f.Extern {
		builder.WriteString(") extern\n")
		return builder.String()
	}
	builder.WriteString(") {\n")
	l = len(f.Blocks)
	for i, bb := range f.Blocks {
		builder.WriteString(bb.String())
		if i+1 == l {
			break
		}
		builder.WriteString("\n")
	}

	builder.WriteString("}\n")
	return builder.String()
}
