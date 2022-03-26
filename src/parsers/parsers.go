package parsers

import (
    "github.com/galaxia-team/void/src/exception"
    "github.com/galaxia-team/void/src/operators"
    "github.com/galaxia-team/void/src/types"
    "github.com/galaxia-team/void/src/utils"
    "strings"
    "regexp"
    "fmt"
)

func ParseBoolWrapper(b string) bool {
    nb := types.ConvertType(b)
    bt := types.GetType(nb)

    switch (bt) {
        case "bool":
            return nb
        case "int":
            cb, ok := nb.(int)
            return cb > 0
        case "float":
            cb, ok := nb.(float64)
            return cb >= 1
    }

    exception.Except("not_bool", utils.CurrentLine)
    return
}

func ParseArray(a string, s byte) interface{} {
    var e byte

    switch (s) {
        case "[":
            e = "]"
        case "(":
            e = ")"
        default:
            fmt.Println("start char: ", s)
            panic("problem in ParseArray")
    }

    restr1 := regexp.MustCompile(`".*"`)
    restr2 := regexp.MustCompile(`'.*'`)
    
    strs := utils.Extend(
        restr1.FindAllString(a, -1),
        restr2.FindAllString(a, -1)
    )
    
    for n, str := range strs {
        a = strings.Replace(a, str, fmt.Sprintf("VOID_PA_STR_%f", n), -1)
    }

    return a
}

func ParseStatement(s string) interface{} {
    return s
}