package operators

import (
    "github.com/galaxia-team/void/void/src/exception"
    "github.com/galaxia-team/void/void/src/types"
    "math"
    "fmt"
)

var (
    StringOps = map[string]func(string, string) string {
        "+": AddS,
    }
    IntOps = map[string]func(int, int) int {
        "+": AddI,
        "-": SubI,
        "*": MulI,
        "**": PowI,
        "/": DivI,
        "//": FloorDivI,
        "%": ModI,
    }
    FloatOps = map[string]func(float64, float64) float64 {
        "+": AddF,
        "-": SubF,
        "*": MulF,
        "**": PowF,
        "/": DivF,
        "//": FloorDivF,
    }
    BoolOps = map[string]func(bool, bool) bool {
        "&&": And,
        "||": Or,
    }
)

func ApplyOperator(op string, x interface{}, y interface{}, n int) interface{} {
    xt := types.GetType(x)
    yt := types.GetType(y)

    var xok bool
    var yok bool
    
    var sx string
    var ix int
    var fx float64
    var bx bool

    var sy string
    var iy int
    var fy float64
    var by bool
    
    switch (xt) {
        case "string":
            sx, xok = x.(string)
        case "int":
            ix, xok = x.(int)
        case "float":
            fx, xok = x.(float64)
        case "bool":
            bx, xok = x.(bool)
    }

    switch (yt) {
        case "string":
            sy, yok = y.(string)
        case "int":
            iy, yok = y.(int)
        case "float":
            fy, yok = y.(float64)
        case "bool":
            by, yok = y.(bool)
    }
    
    if !xok || !yok {
        exception.Except("invalid_type", n)
    }
    
    if _, ok := BoolOps[op]; !ok {
        if xt == "string" && yt == "string" {
            if _, ok := StringOps[op]; ok {
                StringOps[op](sx, sy)
            } else {
                exception.Except("invalid_op", n)
            }
        } else if yt == "int" && xt == "int" {
            if _, ok := IntOps[op]; ok {
                IntOps[op](ix, iy)
            } else {
                exception.Except("invalid_op", n)
            }
        } else if xt == "int" && yt == "float" {
            fx = float64(ix)
            if _, ok := FloatOps[op]; ok {
                return FloatOps[op](fx, fy)
            } else {
                exception.Except("invalid_op", n)
            }
        } else if xt == "float" && yt == "int" {
            fy = float64(iy)
            if _, ok := FloatOps[op]; ok {
                return FloatOps[op](fx, fy)
            } else {
                exception.Except("invalid_op", n)
            }
        } else if xt == "float" && yt == "float" {
            if _, ok := FloatOps[op]; ok {
                return FloatOps[op](fx, fy)
            } else {
                exception.Except("invalid_op", n)
            }
        } else {
            exception.Except("invalid_type", n)
        }
    } else {
        if sx != "" {
            bx = types.GetBool(sx, n)
        } else if ix != 0 {
            bx = types.GetBool(fmt.Sprintf("%d", ix), n)
        } else if fx != 0 {
            bx = types.GetBool(fmt.Sprintf("%f", fx), n)
        } else {
            by = false
        }

        if sy != "" {
            by = types.GetBool(sy, n)
        } else if iy != 0 {
            by = types.GetBool(fmt.Sprintf("%d", iy), n)
        } else if fy != 0 {
            by = types.GetBool(fmt.Sprintf("%f", fy), n)
        } else {
            by = false
        }
        
        return BoolOps[op](bx, by)
    }
    
    return ""
}

func AddI(x int, y int) int {
    return x + y
}

func AddF(x float64, y float64) float64 {
    return x + y
}

func AddS(x string, y string) string {
    return x + y
}

func SubI(x int, y int) int {
    return x - y
}

func SubF(x float64, y float64) float64 {
    return x - y
}

func MulI(x int, y int) int {
    return x * y
}

func MulF(x float64, y float64) float64 {
    return x * y
}

func PowI(x int, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func PowF(x float64, y float64) float64 {
    return math.Pow(x, y)
}

func DivI(x int, y int) int {
    return x / y
}

func DivF(x float64, y float64) float64 {
    return x / y
}

func FloorDivI(x int, y int) int {
    return int(math.Floor(float64(x) / float64(y)))
}

func FloorDivF(x float64, y float64) float64 {
    return math.Floor(x / y)
}

func ModI(x int, y int) int {
    return x % y
}

func And(x bool, y bool) bool {
    return x && y
}

func Or(x bool, y bool) bool {
    return x || y
}
