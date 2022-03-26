package variables

import (
    "github.com/galaxia-team/void/src/exception"
    "github.com/galaxia-team/void/src/parsers"
    "github.com/galaxia-team/void/src/utils"
    "github.com/galaxia-team/void/src/types"
    "strings"
)

type Var struct {
    mutable bool
    name, scope, val string
}

var LocalVars = map[string]map[string]Var {}

var GlobalVars = map[string]Var {}

func InitVar(fc []string, vs string, n int) {
    vt := fc[0]
    vn := fc[1]
    if fc[2] != "=" {
        exception.Except("invalid_syntax", n)
    }
    fc = utils.RemoveIndex(fc, 0)
    fc = utils.RemoveIndex(fc, 0)
    fc = utils.RemoveIndex(fc, 0)
    vst := strings.Join(fc, " ")
    vv := parsers.ParseStatement(vst)
    if vt == "let" {
        if vs != "root" {
            LocalVars[vs][vn] = Var{
                mutable: true,
                name: vn,
                scope: vs,
                val: vv,
            }
        } else {
            GlobalVars[vn] = Var{
                mutable: true,
                name: vn,
                scope: vs,
                val: vv,
            }
        }
    } else if _, ok := Vars[vn]; !ok {
        if vs != "root" {
            LocalVars[vs][vn] = Var{
                mutable: false,
                name: vn,
                scope: vs,
                val: vv,
            }
        } else {
            GlobalVars[vn] = Var{
                mutable: false,
                name: vn,
                scope: vs,
                val: vv,
            }
        }     
    } else {
        exception.Except("const_immutable", n)
    }
}