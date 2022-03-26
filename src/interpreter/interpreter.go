package interpreter

import (
    //"github.com/galaxia-team/void/src/exception"
    "github.com/galaxia-team/void/src/variables"
    "github.com/galaxia-team/void/src/scopes"
    "github.com/galaxia-team/void/src/utils"
    "strings"
    //"fmt"
)

type Arg struct {
    name, val string
}

func Interpret(d []string, a []Arg) {
    for cn, cc := range d {
        s := scopes.GetScope(cn, cc)
        fc := strings.Fields(cc)
        if utils.Contains([]string{"let", "const"}, fc[0]) {
            variables.InitVar(fc, s, cn)
        }
    }
}