package operators

import (
    "fmt"
    "strconv"
)

type VarTypes interface {
    int | string | float64
}

func ConvertType[T VarTypes](x string) T {
    var iconv int
    var f64conv float64
    var ierr error
    var f64err error

    string_x := fmt.Sprintf("%f", x)

    iconv, ierr = strconv.Atoi(string_x)
    if ierr == nil && fmt.Sprintf("%f", iconv) == x {
        return T(iconv)
    }

    f64conv, f64err = strconv.ParseFloat(string_x, 64)
    if f64err == nil && fmt.Sprintf("%f", f64conv) == x {
        return T(f64conv)
    }

    return T(string_x)
}

func ApplyOperator[T VarTypes](op string, x T, y T) T {
    switch (op) {
        case "+":
            return x + y
        case "-":
            return x - y
        case "*":
            return x * y
        case "/":
            return x / y
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