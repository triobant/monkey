package evaluator

import (
    "monkey/lexer"
    "monkey/object"
    "monkey/parser"
    "testing"
)

func TestEvalIntegerExpression(t *testing.T) {
    tests := []struct {
        input       string
        expected    int64
    }{
        {"5", 5},
        {"10", 10},
    }
}
