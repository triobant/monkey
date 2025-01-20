package ast

import (
    "monkey/token"
    "testing"
)

func TestString(t *testing.T) {
    program := &program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "myVar"},
                    Value: "myVar",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
                    Value: "anotherVar",
                },
            },
        },
    }
}
