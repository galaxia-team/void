package lib

import (
    "strings"
)

func ImportLib(l string) {
    i = strings.Replace(l, ".", "/", -1)
    import il "github.com/galaxia-team/void/src/lib/"+i
    return il
}
