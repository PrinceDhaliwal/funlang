package types

import "strings"

// represents structs
type structType struct {
    elems map[string]Type
}

func (s *structType) Elem() Type {
    return nil
}

func (s *structType) Field(name string) Type {
    t, ok := s.elems[name]
    if !ok {
        return nil
    }

    return t
}

func (s *structType) Tag() TypeTag {
    return STRUCT_TYPE
}

func (s *structType) Name() string {
    builder := strings.Builder{}
    builder.WriteString("struct{")
    l := len(s.elems)
    i := 0;
    for name, t := range s.elems {
        builder.WriteString(name)
        builder.WriteString(":")
        builder.WriteString(t.Name())

        if i+1 != l {
            builder.WriteString(";")
        }
        i++
    }

    builder.WriteString("}")
    return builder.String()
}


func (s *structType) String() string {
    return s.Name()
}

func (s *structType) LenFields() int {
    return len(s.elems)
}

func ToStructType(t Type) *structType {
    if t.Tag() != STRUCT_TYPE {
        return nil
    }

    return t.(*structType)
}
