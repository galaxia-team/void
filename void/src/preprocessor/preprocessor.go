package preprocessor

import (
    "github.com/galaxia-team/void/void/src/interpreter"
    "github.com/galaxia-team/void/void/src/utils"
    "strings"
)

func PreProcess(s []string) {
    for n, cl := range s {
        s[n] = strings.TrimSpace(cl)
        /*
        if s[n][0:8] == "include " {
            ti = strings.Replace(s[n], "include ", "", -1)
        }
        */
        if s[n] != "" {
            if s[n][0:2] == "//" {
                s[n] = ""
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
    }
    for utils.Contains(s, "") {
        s = utils.Remove(s, "")
    }
    interpreter.Interpret(s, []interpreter.Arg{})
}
