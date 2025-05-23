package evaluator

import (
    "monkey/object"
)

var builtins = map[string]*object.Builtin{
    "len": &object.Builtin{
        Fn: func(args ..object.Object) object.Object {
            if len(args) != 1 {
                return newError("wrong number of arguments. got=%d, want=1",
                    len(args))
            }
        },
    },
}
