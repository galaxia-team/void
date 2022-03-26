package parsers

import (
    //"github.com/galaxia-team/void/src/exception"
    //"github.com/galaxia-team/void/src/operators"
    //"github.com/galaxia-team/void/src/types"
    "github.com/galaxia-team/void/src/utils"
    "strings"
    "regexp"
    "fmt"
)

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