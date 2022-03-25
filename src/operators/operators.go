package operators

import (
    "github.com/galaxia-team/void/src/exception"
    "github.com/galaxia-team/void/src/types"
    "strconv"
    "reflect"
    "math"
    "fmt"
)

const LogicalOps = []string{"&&", "||"}

func ApplyOperator(op string, x interface{}, y interface{}) interface{} {
    xst = types.GetInterfaceType(x)
    yst = types.GetInterfaceType(y)

    if v, ok := x.(string); ok {
        x = v
    } else if v, ok := x.(int); ok {
        x = v
    } else if v, ok := x.(uint); ok {
        x = v
    } else if v, ok := x.(float64); ok {
        x = v
    } else if v, ok := x.(bool); ok {
        x = v
    }

    if v, ok := y.(string); ok {
        y = v
    } else if v, ok := y.(int); ok {
        y = v
    } else if v, ok := y.(uint); ok {
        y = v
    } else if v, ok := y.(float64); ok {
        y = v
    } else if v, ok := y.(bool); ok {
        y = v
    }

    xt = reflect.TypeOf(x)
    yt = reflect.TypeOf(y)

    if xt == string && yt == string {
        switch (op) {
            case "+":
                return x + y, nil
            default:
                return "", errors.New("invalid operation")
        }
    } else if xt == int && yt == int {
        switch (op) {
            case "+":
                return x + y, nil
            default:
                return "", errors.New("invalid operation")
        }
    }

    switch (op) {
        case "+":
            return x + y
        case "-":
            return x - y
        case "*":
            return x * y
        case "**":
            return math.Pow(x, y)
        case "/":
            return x / y
        case "//":
            return math.Floor(x / y)
        case "%":
            return x % y
        case "&":
            return x & y
        case "|":
            return x | y
        case "^":
            return x ^ y
        case "&^":
            return x &^ y
        case "<<":
            return x << y
        case ">>":
            return x >> y
    }
}