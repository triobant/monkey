package parser

import (
    "fmt"
    "strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
    return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
    fmt.Printf("%s%s\n", identLevel(), fs)
}
