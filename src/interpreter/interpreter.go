package interpreter

import (
    /*
    "github.com/galaxia-team/void/src/exception"
    "github.com/galaxia-team/void/src/variables"
    */
    "fmt"
)

type Arg struct {
    name, val string
}

func Interpret(data []string, args []Arg) {
    for lnum, line := range data {
        fmt.Println(line)
        fmt.Println(lnum)
    }
}