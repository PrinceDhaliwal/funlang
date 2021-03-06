// literal.go defines the literal types
package types

// represents an integer
type intType struct {
}

func (t *intType) Elem() Type {
    return nil
}

func (t *intType) Field(string) Type {
    return nil
}

func (t *intType) Tag() TypeTag {
    return INT_TYPE
}

func (t *intType) Name() string {
    return "int"
}

func (t *intType) String() string {
    return t.Name()
}

type floatType struct {}

func (t *floatType) String() string {
    return t.Name()
}

func (t *floatType) Elem() Type {
    return nil
}

func (t *floatType) Field(string) Type {
    return nil
}

func (t *floatType) Tag() TypeTag {
    return FLOAT_TYPE
}

func (t *floatType) Name() string {
    return "float"
}

type stringType struct {}

func (t *stringType) String() string {
    return t.Name()
}
func (t *stringType) Name() string {
    return "string"
}

func (t *stringType) Elem() Type {
    return nil
}

func (t *stringType) Field(string) Type {
    return nil
}

func (t *stringType) Tag() TypeTag {
    return STRING_TYPE
}

func ToIntType(t Type) *intType {
    return t.(*intType)
}

func ToFloatType(t Type) *floatType {
    return t.(*floatType)
}

func ToStringType(t Type) *stringType {
    return t.(*stringType)
}
