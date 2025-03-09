package object

import "fmt"

type ObjectType string

const (
    INTEGER_OBJ = "INTEGER"
)

type Object interface {
    Type()      ObjectType
    Inspect()   String
}

type Integer struct {
    Value       int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
