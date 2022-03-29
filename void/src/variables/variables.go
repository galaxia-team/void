package variables

import (
    //"github.com/galaxia-team/void/void/src/exception"
    //"github.com/galaxia-team/void/void/src/parsers"
    //"github.com/galaxia-team/void/void/src/utils"
    //"github.com/galaxia-team/void/void/src/types"
    //"strings"
)

func InitVar(fc []string, vs string, n int) {
    return
}

/*
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
            utils.LocalVars[vs][vn] = utils.Var {
                mutable: true,
                name: vn,
                scope: vs,
                val: vv,
            }
        } else {
            utils.GlobalVars[vn] = utils.Var {
                mutable: true,
                name: vn,
                scope: vs,
                val: vv,
            }
        }
    } else if _, ok := Vars[vn]; !ok {
        if vs != "root" {
            utils.LocalVars[vs][vn] = utils.Var {
                mutable: false,
                name: vn,
                scope: vs,
                val: vv,
            }
        } else {
            utils.GlobalVars[vn] = utils.Var {
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
*/
