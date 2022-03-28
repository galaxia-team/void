package types

import (
    "github.com/galaxia-team/void/src/exception"
    "strconv"
    "fmt"
)

func GetBool(b string, n int) bool {
    nb := ConvertType(b)
    bt := GetType(nb)

    switch (bt) {
        case "string":
            sb, _ := nb.(string)
            return sb != ""
        case "bool":
            bb, _ := nb.(bool)
            return bb
        case "int":
            cb, _ := nb.(int)
            return cb > 0
        case "float":
            cb, _ := nb.(float64)
            return cb >= 1
    }

    exception.Except("not_bool", n)
    return false
}

func GetType(x interface{}) string {
    if _, ok := x.(int); ok {
        return "int"
    } else if _, ok := x.(float64); ok {
        return "float"
    } else if _, ok := x.(bool); ok {
        return "bool"
    }
    return "string"
}

func ConvertType(x string) interface{} {
    var ierr error
    var i64conv int64
    
    if x[0:2] == "0x" {
        i64conv, ierr = strconv.ParseInt(x, 16, 64)
    } else {
        i64conv, ierr = strconv.ParseInt(x, 10, 64)
    }
    
    iconv := int(i64conv)
    
    if ierr == nil && fmt.Sprintf("%f", iconv) == x {
        return iconv
    }

    fconv, ferr := strconv.ParseFloat(x, 64)
    if ferr == nil && fmt.Sprintf("%f", fconv) == x {
        return fconv
    }

    bconv, berr := strconv.ParseBool(x)
    if berr == nil && fmt.Sprintf("%f", bconv) == x {
        return bconv
    }

    return x
}
