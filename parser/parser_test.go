package parser

import (
    "testing"
    "monkey/ast"
    "monkey/lexer"
)

func TestLetStatements(t *testing.T) {
    input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
}
