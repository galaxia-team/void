package parsers

import (
    //"github.com/galaxia-team/void/void/src/exception"
    //"github.com/galaxia-team/void/void/src/operators"
    //"github.com/galaxia-team/void/void/src/types"
    "github.com/galaxia-team/void/void/src/utils"
    "strings"
    "regexp"
    "fmt"
)

func ParseLiteral(l string, n int) interface{} {
    return l
}

func ParseArray(a string, s string) interface{} {
    var e string
    
    fmt.Println("start char: ", s)
    
    switch (s) {
        case "[":
            e = "]"
        case "(":
            e = ")"
        default:
            panic("problem in ParseArray")
    }
    
    fmt.Println("end char: ", e)

    refstr1 := regexp.MustCompile(`f".*"`)
    refstr2 := regexp.MustCompile(`f'.*'`)
    
    fstrs := utils.Extend(refstr1.FindAllString(a, -1), refstr2.FindAllString(a, -1))

    for fn, fstr := range fstrs {
        a = strings.Replace(a, fstr, fmt.Sprintf("VOID_PA_FSTR_%f", fn), -1)
    }

    restr1 := regexp.MustCompile(`".*"`)
    restr2 := regexp.MustCompile(`'.*'`)
    
    strs := utils.Extend(restr1.FindAllString(a, -1), restr2.FindAllString(a, -1))
    
    for n, str := range strs {
        a = strings.Replace(a, str, fmt.Sprintf("VOID_PA_STR_%f", n), -1)
    }

    return a
}

func ParseStatement(s string) interface{} {
    return s
}
