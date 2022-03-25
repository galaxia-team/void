package preprocessor

import (
    "github.com/galaxia-team/void/src/interpreter"
    "github.com/galaxia-team/void/src/utils"
    "strings"
)

func PreProcess(s []string) {
    for n, cl := range s {
        s[n] = strings.TrimSpace(cl)
        if s[n][0:8] == "include " {
            ti = 
        }
        if s[n][0:2] == "//" {
            s[n] == ""
        }
        if s[n][0:2] == "/*" {
            ccn := n
            ccl := cl
            for ccl[len(cl)-2:] == "*/" {
                ccn += 1
                ccl = s[ccn]
                s[ccn] = ""
            }
        }
    }
    for utils.Contains(s, "") {
        s = utils.Remove(s, "")
    }
    interpreter.Interpret(s, []interpreter.Arg{})
}