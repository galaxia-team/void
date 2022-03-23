package preprocessor

import (
    "github.com/galaxia-team/void/src/interpreter"
    "github.com/galaxia-team/void/src/utils"
)

func PreProcess(s []string) {
    for utils.Contains(s, "") {
        s = utils.Remove(s, "")
    }
    interpreter.Interpret(s, []interpreter.Arg{})
}