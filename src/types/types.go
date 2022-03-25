package types

import (
    "fmt"
    "strconv"
)

const TrueStrings = []string{"true", "True", "TRUE"}

const FalseStrings = []string{"false", "False", "FALSE"}

func GetInterfaceType(x interface{}) string {
    if _, ok := x.(string); ok {
        return "string"
    } else if _, ok := x.(int); ok {
        return "int"
    } else if _, ok := x.(uint); ok {
        return "uint"
    } else if _, ok := x.(float64); ok {
        return "float"
    } else if _, ok := x.(bool); ok {
        return "bool"
    } else {
        panic("OHHH SHI- (GetInterfaceType issue)")
    }
}

func PsuedoBool(pb interface{}) bool {
    
    if pb[0:1] == "!" {
        
    }
}

func ParseType(x string) interface{} {
    iconv, ierr := strconv.ParseInt(x)
    if ierr == nil && fmt.Sprintf("%f", iconv) == x {
        return iconv
    }

    fconv, ferr := strconv.ParseFloat(x, 64)
    if ferr == nil && fmt.Sprintf("%f", fconv) == x {
        return fconv
    }

    uiconv, uierr := strconv.ParseUint(x)
    if uierr == nil && fmt.Sprintf("%f", uiconv) == x {
        return uiconv
    }

    bconv, berr := strconv.ParseBool(x)
    if berr == nil && fmt.Sprintf("%f", bconv) == x {
        return bconv
    }

    return x
}