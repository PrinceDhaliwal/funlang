package ssa

import (
	"bitbucket.org/dhaliwalprince/funlang/types"
	"fmt"
)

type ConstantInt struct {
	valueWithNoName
	valueWithUsers
	Value int
}

func (c *ConstantInt) Uses() []Value {
	return []Value{}
}

func (c *ConstantInt) String() string {
	return fmt.Sprintf("%d:%s", c.Value, c.Type())
}

func (c *ConstantInt) ShortString() string {
	return c.String()
}

func (c *ConstantInt) Tag() ValueTag {
	return CONSTANT_INT
}

func (c *ConstantInt) Type() types.Type {
	return typeFactory.IntType()
}

type ConstantString struct {
	valueWithNoName
	valueWithUsers
	Value string
}

func (c *ConstantString) String() string {
	return c.Value+":"+fmt.Sprint(c.Type())
}

func (c *ConstantString) ShortString() string {
	return c.String()
}

func (c *ConstantString) Uses() []Value {
	return []Value{}
}

func (c *ConstantString) Tag() ValueTag {
	return CONSTANT_STRING
}

func (c *ConstantString) Type() types.Type {
	return typeFactory.StringType()
}
