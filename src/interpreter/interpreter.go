package interpreter

import (
    //"github.com/galaxia-team/void/src/exception"
    "github.com/galaxia-team/void/src/variables"
    "github.com/galaxia-team/void/src/utils"
    "strings"
    "fmt"
)

type Arg struct {
    name, val string
}

func Interpret(d []string, a []Arg) {
    for cn, cc := range d {
        utils.CurrentLine = cn
        fc := strings.Fields(cc)
        if utils.Contains([]string{"let", "const"}, fc[0]) {
            variables.InitVar(fc, scope, cn)
        }
    }
}