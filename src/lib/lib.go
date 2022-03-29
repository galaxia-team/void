package lib

import (
    "github.com/galaxia-team/void/src/lib/console"
    "strings"
)

func ImportLib(l string) {
    i = strings.Replace(l, ".", "/", -1)
    import "github.com/galaxia-team/void/src/lib/"+i
    return
}
