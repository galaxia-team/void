package parsers

import (
    "github.com/galaxia-team/void/src/operators"
    "github.com/galaxia-team/void/src/types"
    "regexp"
)

func ParseArray(a string, s byte) interface{} {
    var e byte

    switch (s) {
        case "[":
            e = "]"
        case "(":
            e = ")"
    }

    restr1 := regexp.MustCompile(`".*"`)
    restr2 := regexp.MustCompile(`'.*'`)

    dqstrs := restr1.FindAllString(a, -1)
    sqstrs := restr2.FindAllString(a, -1)

    return a
}

func ParseStatement(s string) interface{} {
    return s
}