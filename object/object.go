package object

import "fmt"

type ObjectType string

type Object interface {
    Type()      ObjectType
    Inspect()   String
}

type Integer struct {
    Value       int64
}
