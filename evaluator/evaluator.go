package evaluator

import (
    "monkey/ast"
    "monkey/object"
)

func Eval(node ast.Node) object.Object {
    switch node := node.(type) {

    // Statements
    case *ast.Program:
        return evalStatements(node.Statements)

    case *ast.ExpressionStatement:
        return Eval(node.Expression)

    // Expressions
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}

    }

    return nil
}
